package com.wandoulabs.jodis.auto.log;


import org.apache.log4j.Level;
import org.apache.log4j.Logger;
import org.productivity.java.syslog4j.impl.log4j.Syslog4jAppender;

import java.util.concurrent.ArrayBlockingQueue;
import java.util.concurrent.Executors;
import java.util.concurrent.ThreadPoolExecutor;
import java.util.concurrent.TimeUnit;

/**
 * Created by zhufeng on 14-10-20.
 */
public class LogHandler {

    private static String host;
    private static String port;
    private static String appName;

    public static Syslog4jAppender remoteAppender ;

    public static Logger rlog = Logger.getLogger("log-client-remote");

    /**
     * 初始化线程池 ONE
     * corePoolSize 10
     * maxPoolSize 10
     * keepLiveTime 60 second
     * 有边界的queue
     * 超出最大线程数则将积压任务交给主线程处理
     */
    public static ThreadPoolExecutor EXECUTOR_ONE = new ThreadPoolExecutor(
            5,
            5,
            60,
            TimeUnit.SECONDS,
            new ArrayBlockingQueue<Runnable>(5000),
            Executors.defaultThreadFactory(),
            new ThreadPoolExecutor.CallerRunsPolicy()
    );


    static {
            remoteAppender = new Syslog4jAppender();
            remoteAppender.setPort(port);
            remoteAppender.setProtocol("udp");
            remoteAppender.setSyslogHost(host);
            remoteAppender.setFacility("local7");
            remoteAppender.activateOptions();
            //初始化
            rlog.removeAllAppenders();
            //设置输出级别
            rlog.setLevel(Level.INFO);
            //是否继承父Logger
            rlog.setAdditivity(false);
            rlog.addAppender(remoteAppender);
    }


    /**
     * 输出到远程syslog
     * @param info
     */
    public static void printFlumeLog(String info){
        EXECUTOR_ONE.submit(new RemoteThread(rlog,info));
    }


    public static void setHost(String host) {
        LogHandler.host = host;
    }

    public static void setPort(String port) {
        LogHandler.port = port;
    }

    public static void setAppName(String appName) {
        LogHandler.appName = appName;
    }

    public static String getHost() {
        return host;
    }

    public static String getPort() {
        return port;
    }

    public static String getAppName() {
        return appName;
    }
}
