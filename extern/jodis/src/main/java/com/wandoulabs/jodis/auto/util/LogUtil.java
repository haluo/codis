package com.wandoulabs.jodis.auto.util;

import java.io.ByteArrayOutputStream;
import java.io.IOException;
import java.io.ObjectOutputStream;

/**
 * Created by zhufeng on 15/9/10.
 */
public class LogUtil {
    private static MemoryCounter  mc = new MemoryCounter();


    public static long  objectSize(Object[] os){
        long i = 0;
        for(Object o:os){
            //判断类型 是否为字符串
            if (o instanceof String){
                long bs = String.valueOf(o.toString()).getBytes().length;
                i+=bs;
            }else {
                long bs = mc.estimate(o);
                i += bs;
            }
        }
        return i;
    }

    public static String  strL(Object[] os){
        String s = "";
        for(Object o:os){
            //判断类型 是否为字符串
            if (o instanceof String){
                long bs = o.toString().getBytes().length;
                s+=String.valueOf(o.toString())+":"+bs+"|";
            }
        }
        if(StringUtil.isBlank(s)){
            s = "-";
        }else {
            s = StringUtil.substringBeforeLast(s,"|");
        }
        return s;
    }

}
