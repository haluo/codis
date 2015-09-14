package com.wandoulabs.jodis.auto;

import com.wandoulabs.jodis.auto.log.LogHandler;
import com.wandoulabs.jodis.auto.util.LogUtil;
import com.wandoulabs.jodis.auto.util.TimeUtil;

import java.lang.reflect.InvocationHandler;
import java.lang.reflect.Method;
import java.lang.reflect.Proxy;
import java.util.Date;

/**
 * Created by zhufeng on 15/9/9.
 *
 * 代理类  添加日志信息
 */
public class AutoJedisProxy implements InvocationHandler {
    private Object target;
    public Object bind(Object target) {
        this.target = target;
        return Proxy.newProxyInstance(target.getClass().getClassLoader(),
                target.getClass().getInterfaces(), this);
    }

    @Override
    public Object invoke(Object proxy, Method method, Object[] args) throws Throwable {
        Object result=null;
        long begin = System.currentTimeMillis();
        result=method.invoke(target, args);

        try {
            long cost = System.currentTimeMillis()-begin;
            long size = LogUtil.objectSize(args);
            String sl = LogUtil.strL(args);
            String log = "`"+ TimeUtil.timeForStr(new Date())+"`"+ LogHandler.getAppName()+"`"+method.getName()+"`"+size+"`"+sl+"`"+cost+"`";
            LogHandler.printFlumeLog(log);
        } catch (Throwable t) {
        }
        return result;
    }

}
