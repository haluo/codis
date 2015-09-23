package com.wandoulabs.jodis;

import com.wandoulabs.jodis.auto.AutoJedis;
import redis.clients.jedis.Jedis;
import redis.clients.jedis.JedisPoolConfig;

/**
 * Created by zhufeng on 15/9/8.
 */
public class Test {
    public static void main(String args[]){
        JedisPoolConfig config = new JedisPoolConfig();
        config.setMaxTotal(100);
        config.setMaxIdle(8);
        config.setMaxWaitMillis(1000 * 10);
        config.setTestOnBorrow(false);

        //多个zkserver用逗号分隔，eg：zkserver1:port,zkserver2:prot
        JedisResourcePool jedisPool = new RoundRobinJedisPool("codis_local_123:2181,codis_local_125:2181", 30000, "/zk/codis/db_auto_test/proxy", config);
        jedisPool.register("com.autohome.order.center","192.168.252.44","6008");
        AutoJedis jedis = jedisPool.getAutoResource();
        jedis.set("foo","hahaha");
        String value = jedis.get("foo");
        System.out.println(value);
        //We do not have a returnResource method, just close the jedis instance
        jedis.close();
    }
}
