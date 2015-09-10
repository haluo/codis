package com.wandoulabs.jodis.auto;

/**
 * Created by zhufeng on 15/9/9.
 */
public interface AutoJedis extends redis.clients.jedis.JedisCommands,
        redis.clients.jedis.MultiKeyCommands, redis.clients.jedis.AdvancedJedisCommands, redis.clients.jedis.ScriptingCommands,
        redis.clients.jedis.BasicCommands, redis.clients.jedis.ClusterCommands, redis.clients.jedis.SentinelCommands,
        redis.clients.jedis.BinaryJedisCommands, redis.clients.jedis.MultiKeyBinaryCommands, redis.clients.jedis.AdvancedBinaryJedisCommands,
        redis.clients.jedis.BinaryScriptingCommands, java.io.Closeable  {

    public void close();

}
