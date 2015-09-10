package com.wandoulabs.jodis.auto.log;

import org.apache.log4j.Logger;

/**
 * Created by zhufeng on 14-11-5.
 */
public class RemoteThread implements Runnable{
    private Logger rlog;
    private String info;

    public RemoteThread(Logger rlog, String info) {
        this.rlog = rlog;
        this.info = info;
    }
    public void run() {
        rlog.warn(info);
    }
}
