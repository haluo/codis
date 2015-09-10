package com.wandoulabs.jodis.auto.util;

import java.text.SimpleDateFormat;
import java.util.Date;

/**
 * Created by zhufeng on 14-10-19.
 */
public class TimeUtil {
    private static SimpleDateFormat sf = new SimpleDateFormat("yyyy-MM-dd HH:mm:ss");

    public static String timeForStr(Date date){
        return sf.format(date);
    }
}
