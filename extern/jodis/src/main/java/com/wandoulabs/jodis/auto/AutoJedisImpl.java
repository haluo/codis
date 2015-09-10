package com.wandoulabs.jodis.auto;

import redis.clients.jedis.*;
import redis.clients.util.Slowlog;


import java.util.Collection;
import java.util.List;
import java.util.Map;
import java.util.Set;

/**
 * Created by zhufeng on 15/9/9.
 */
public class AutoJedisImpl implements AutoJedis {
    private Jedis jedis;

    public AutoJedisImpl(Jedis jedis) {
        this.jedis = jedis;
    }

    @Override
    public List<byte[]> configGet(byte[] bytes) {
        return jedis.configGet(bytes);
    }

    @Override
    public byte[] configSet(byte[] bytes, byte[] bytes2) {
        return jedis.configSet(bytes,bytes2);
    }

    @Override
    public List<byte[]> slowlogGetBinary() {
        return jedis.slowlogGetBinary();
    }

    @Override
    public List<byte[]> slowlogGetBinary(long l) {
        return jedis.slowlogGetBinary(l);
    }

    @Override
    public Long objectRefcount(byte[] bytes) {
        return jedis.objectRefcount(bytes);
    }

    @Override
    public byte[] objectEncoding(byte[] bytes) {
        return jedis.objectEncoding(bytes);
    }

    @Override
    public Long objectIdletime(byte[] bytes) {
        return jedis.objectIdletime(bytes);
    }

    @Override
    public String slowlogReset() {
        return jedis.slowlogReset();
    }

    @Override
    public Long slowlogLen() {
        return jedis.slowlogLen();
    }

    @Override
    public String ping() {
        return jedis.ping();
    }

    @Override
    public String quit() {
        return jedis.quit();
    }

    @Override
    public String flushDB() {
        return jedis.flushDB();
    }

    @Override
    public Long dbSize() {
        return jedis.dbSize();
    }

    @Override
    public String select(int i) {
        return jedis.select(i);
    }

    @Override
    public String flushAll() {
        return jedis.flushAll();
    }

    @Override
    public String auth(String s) {
        return jedis.auth(s);
    }

    @Override
    public String save() {
        return jedis.save();
    }

    @Override
    public String bgsave() {
        return jedis.bgsave();
    }

    @Override
    public String bgrewriteaof() {
        return jedis.bgrewriteaof();
    }

    @Override
    public Long lastsave() {
        return jedis.lastsave();
    }

    @Override
    public String shutdown() {
        return jedis.shutdown();
    }

    @Override
    public String info() {
        return jedis.info();
    }

    @Override
    public String info(String s) {
        return jedis.info(s);
    }

    @Override
    public String slaveof(String s, int i) {
        return jedis.slaveof(s,i);
    }

    @Override
    public String slaveofNoOne() {
        return jedis.slaveofNoOne();
    }

    @Override
    public Long getDB() {
        return jedis.getDB();
    }

    @Override
    public String debug(DebugParams debugParams) {
        return jedis.debug(debugParams);
    }

    @Override
    public String configResetStat() {
        return jedis.configResetStat();
    }

    @Override
    public Long waitReplicas(int i, long l) {
        return jedis.waitReplicas(i,l);
    }

    @Override
    public String set(byte[] bytes, byte[] bytes2) {
        return jedis.set(bytes,bytes2);
    }

    @Override
    public byte[] get(byte[] bytes) {
        return jedis.get(bytes);
    }

    @Override
    public Boolean exists(byte[] bytes) {
        return jedis.exists(bytes);
    }

    @Override
    public Long persist(byte[] bytes) {
        return jedis.persist(bytes);
    }

    @Override
    public String type(byte[] bytes) {
        return jedis.type(bytes);
    }

    @Override
    public Long expire(byte[] bytes, int i) {
        return jedis.expire(bytes,i);
    }

    @Override
    public Long expireAt(byte[] bytes, long l) {
        return jedis.expireAt(bytes,l);
    }

    @Override
    public Long ttl(byte[] bytes) {
        return jedis.ttl(bytes);
    }

    @Override
    public Boolean setbit(byte[] bytes, long l, boolean b) {
        return jedis.setbit(bytes,l,b);
    }

    @Override
    public Boolean setbit(byte[] bytes, long l, byte[] bytes2) {
        return jedis.setbit(bytes,l,bytes2);
    }

    @Override
    public Boolean getbit(byte[] bytes, long l) {
        return jedis.getbit(bytes,l);
    }

    @Override
    public Long setrange(byte[] bytes, long l, byte[] bytes2) {
        return jedis.setrange(bytes,l,bytes2);
    }

    @Override
    public byte[] getrange(byte[] bytes, long l, long l2) {
        return jedis.getrange(bytes,l,l2);
    }

    @Override
    public byte[] getSet(byte[] bytes, byte[] bytes2) {
        return jedis.getSet(bytes,bytes2);
    }

    @Override
    public Long setnx(byte[] bytes, byte[] bytes2) {
        return jedis.setnx(bytes,bytes2);
    }

    @Override
    public String setex(byte[] bytes, int i, byte[] bytes2) {
        return jedis.setex(bytes,i,bytes2);
    }

    @Override
    public Long decrBy(byte[] bytes, long l) {
        return jedis.decrBy(bytes,l);
    }

    @Override
    public Long decr(byte[] bytes) {
        return jedis.decr(bytes);
    }

    @Override
    public Long incrBy(byte[] bytes, long l) {
        return jedis.incrBy(bytes,l);
    }

    @Override
    public Double incrByFloat(byte[] bytes, double v) {
        return jedis.incrByFloat(bytes,v);
    }

    @Override
    public Long incr(byte[] bytes) {
        return jedis.incr(bytes);
    }

    @Override
    public Long append(byte[] bytes, byte[] bytes2) {
        return jedis.append(bytes,bytes2);
    }

    @Override
    public byte[] substr(byte[] bytes, int i, int i2) {
        return jedis.substr(bytes,i,i2);
    }

    @Override
    public Long hset(byte[] bytes, byte[] bytes2, byte[] bytes3) {
        return jedis.hset(bytes,bytes2,bytes3);
    }

    @Override
    public byte[] hget(byte[] bytes, byte[] bytes2) {
        return jedis.hget(bytes,bytes2);
    }

    @Override
    public Long hsetnx(byte[] bytes, byte[] bytes2, byte[] bytes3) {
        return jedis.hsetnx(bytes,bytes2,bytes3);
    }

    @Override
    public String hmset(byte[] bytes, Map<byte[], byte[]> map) {
        return jedis.hmset(bytes,map);
    }

    @Override
    public List<byte[]> hmget(byte[] bytes, byte[]... bytes2) {
        return jedis.hmget(bytes,bytes2);
    }

    @Override
    public Long hincrBy(byte[] bytes, byte[] bytes2, long l) {
        return jedis.hincrBy(bytes,bytes2,l);
    }

    @Override
    public Double hincrByFloat(byte[] bytes, byte[] bytes2, double v) {
        return jedis.hincrByFloat(bytes,bytes2,v);
    }

    @Override
    public Boolean hexists(byte[] bytes, byte[] bytes2) {
        return jedis.hexists(bytes,bytes2);
    }

    @Override
    public Long hdel(byte[] bytes, byte[]... bytes2) {
        return jedis.hdel(bytes,bytes2);
    }

    @Override
    public Long hlen(byte[] bytes) {
        return jedis.hlen(bytes);
    }

    @Override
    public Set<byte[]> hkeys(byte[] bytes) {
        return jedis.hkeys(bytes);
    }

    @Override
    public Collection<byte[]> hvals(byte[] bytes) {
        return jedis.hvals(bytes);
    }

    @Override
    public Map<byte[], byte[]> hgetAll(byte[] bytes) {
        return jedis.hgetAll(bytes);
    }

    @Override
    public Long rpush(byte[] bytes, byte[]... bytes2) {
        return jedis.rpush(bytes,bytes2);
    }

    @Override
    public Long lpush(byte[] bytes, byte[]... bytes2) {
        return jedis.lpush(bytes,bytes2);
    }

    @Override
    public Long llen(byte[] bytes) {
        return jedis.llen(bytes);
    }

    @Override
    public List<byte[]> lrange(byte[] bytes, long l, long l2) {
        return jedis.lrange(bytes,l,l2);
    }

    @Override
    public String ltrim(byte[] bytes, long l, long l2) {
        return jedis.ltrim(bytes,l,l2);
    }

    @Override
    public byte[] lindex(byte[] bytes, long l) {
        return jedis.lindex(bytes,l);
    }

    @Override
    public String lset(byte[] bytes, long l, byte[] bytes2) {
        return jedis.lset(bytes,l,bytes2);
    }

    @Override
    public Long lrem(byte[] bytes, long l, byte[] bytes2) {
        return jedis.lrem(bytes,l,bytes2);
    }

    @Override
    public byte[] lpop(byte[] bytes) {
        return jedis.lpop(bytes);
    }

    @Override
    public byte[] rpop(byte[] bytes) {
        return jedis.rpop(bytes);
    }

    @Override
    public Long sadd(byte[] bytes, byte[]... bytes2) {
        return jedis.sadd(bytes,bytes2);
    }

    @Override
    public Set<byte[]> smembers(byte[] bytes) {
        return jedis.smembers(bytes);
    }

    @Override
    public Long srem(byte[] bytes, byte[]... bytes2) {
        return jedis.srem(bytes,bytes2);
    }

    @Override
    public byte[] spop(byte[] bytes) {
        return jedis.spop(bytes);
    }

    @Override
    public Long scard(byte[] bytes) {
        return jedis.scard(bytes);
    }

    @Override
    public Boolean sismember(byte[] bytes, byte[] bytes2) {
        return jedis.sismember(bytes,bytes2);
    }

    @Override
    public byte[] srandmember(byte[] bytes) {
        return jedis.srandmember(bytes);
    }

    @Override
    public List<byte[]> srandmember(byte[] bytes, int i) {
        return jedis.srandmember(bytes,i);
    }

    @Override
    public Long strlen(byte[] bytes) {
        return jedis.strlen(bytes);
    }

    @Override
    public Long zadd(byte[] bytes, double v, byte[] bytes2) {
        return jedis.zadd(bytes,v,bytes2);
    }

    @Override
    public Long zadd(byte[] bytes, Map<byte[], Double> doubleMap) {
        return jedis.zadd(bytes,doubleMap);
    }

    @Override
    public Set<byte[]> zrange(byte[] bytes, long l, long l2) {
        return jedis.zrange(bytes,l,l2);
    }

    @Override
    public Long zrem(byte[] bytes, byte[]... bytes2) {
        return jedis.zrem(bytes,bytes2);
    }

    @Override
    public Double zincrby(byte[] bytes, double v, byte[] bytes2) {
        return jedis.zincrby(bytes,v,bytes2);
    }

    @Override
    public Long zrank(byte[] bytes, byte[] bytes2) {
        return jedis.zrank(bytes,bytes2);
    }

    @Override
    public Long zrevrank(byte[] bytes, byte[] bytes2) {
        return jedis.zrevrank(bytes,bytes2);
    }

    @Override
    public Set<byte[]> zrevrange(byte[] bytes, long l, long l2) {
        return jedis.zrevrange(bytes,l,l2);
    }

    @Override
    public Set<Tuple> zrangeWithScores(byte[] bytes, long l, long l2) {
        return jedis.zrangeWithScores(bytes,l,l2);
    }

    @Override
    public Set<Tuple> zrevrangeWithScores(byte[] bytes, long l, long l2) {
        return jedis.zrevrangeWithScores(bytes,l,l2);
    }

    @Override
    public Long zcard(byte[] bytes) {
        return jedis.zcard(bytes);
    }

    @Override
    public Double zscore(byte[] bytes, byte[] bytes2) {
        return jedis.zscore(bytes,bytes2);
    }

    @Override
    public List<byte[]> sort(byte[] bytes) {
        return jedis.sort(bytes);
    }

    @Override
    public List<byte[]> sort(byte[] bytes, SortingParams sortingParams) {
        return jedis.sort(bytes,sortingParams);
    }

    @Override
    public Long zcount(byte[] bytes, double v, double v2) {
        return jedis.zcount(bytes,v,v2);
    }

    @Override
    public Long zcount(byte[] bytes, byte[] bytes2, byte[] bytes3) {
        return jedis.zcount(bytes,bytes2,bytes3);
    }

    @Override
    public Set<byte[]> zrangeByScore(byte[] bytes, double v, double v2) {
        return jedis.zrangeByScore(bytes,v,v2);
    }

    @Override
    public Set<byte[]> zrangeByScore(byte[] bytes, byte[] bytes2, byte[] bytes3) {
        return jedis.zrangeByScore(bytes,bytes2,bytes3);
    }

    @Override
    public Set<byte[]> zrevrangeByScore(byte[] bytes, double v, double v2) {
        return jedis.zrevrangeByScore(bytes,v,v2);
    }

    @Override
    public Set<byte[]> zrangeByScore(byte[] bytes, double v, double v2, int i, int i2) {
        return jedis.zrangeByScore(bytes,v,v2,i,i2);
    }

    @Override
    public Set<byte[]> zrevrangeByScore(byte[] bytes, byte[] bytes2, byte[] bytes3) {
        return jedis.zrevrangeByScore(bytes,bytes2,bytes3);
    }

    @Override
    public Set<byte[]> zrangeByScore(byte[] bytes, byte[] bytes2, byte[] bytes3, int i, int i2) {
        return jedis.zrangeByScore(bytes,bytes2,bytes3,i,i2);
    }

    @Override
    public Set<byte[]> zrevrangeByScore(byte[] bytes, double v, double v2, int i, int i2) {
        return jedis.zrevrangeByScore(bytes,v,v2,i,i2);
    }

    @Override
    public Set<Tuple> zrangeByScoreWithScores(byte[] bytes, double v, double v2) {
        return jedis.zrangeByScoreWithScores(bytes,v,v2);
    }

    @Override
    public Set<Tuple> zrevrangeByScoreWithScores(byte[] bytes, double v, double v2) {
        return jedis.zrevrangeByScoreWithScores(bytes,v,v2);
    }

    @Override
    public Set<Tuple> zrangeByScoreWithScores(byte[] bytes, double v, double v2, int i, int i2) {
        return jedis.zrangeByScoreWithScores(bytes,v,v2,i,i2);
    }

    @Override
    public Set<byte[]> zrevrangeByScore(byte[] bytes, byte[] bytes2, byte[] bytes3, int i, int i2) {
        return jedis.zrevrangeByScore(bytes,bytes2,bytes3,i,i2);
    }

    @Override
    public Set<Tuple> zrangeByScoreWithScores(byte[] bytes, byte[] bytes2, byte[] bytes3) {
        return jedis.zrangeByScoreWithScores(bytes,bytes2,bytes3);
    }

    @Override
    public Set<Tuple> zrevrangeByScoreWithScores(byte[] bytes, byte[] bytes2, byte[] bytes3) {
        return jedis.zrevrangeByScoreWithScores(bytes,bytes2,bytes3);
    }

    @Override
    public Set<Tuple> zrangeByScoreWithScores(byte[] bytes, byte[] bytes2, byte[] bytes3, int i, int i2) {
        return jedis.zrangeByScoreWithScores(bytes,bytes2,bytes3,i,i2);
    }

    @Override
    public Set<Tuple> zrevrangeByScoreWithScores(byte[] bytes, double v, double v2, int i, int i2) {
        return jedis.zrevrangeByScoreWithScores(bytes,v,v2,i,i2);
    }

    @Override
    public Set<Tuple> zrevrangeByScoreWithScores(byte[] bytes, byte[] bytes2, byte[] bytes3, int i, int i2) {
        return jedis.zrevrangeByScoreWithScores(bytes,bytes2,bytes3,i,i2);
    }

    @Override
    public Long zremrangeByRank(byte[] bytes, long l, long l2) {
        return jedis.zremrangeByRank(bytes,l,l2);
    }

    @Override
    public Long zremrangeByScore(byte[] bytes, double v, double v2) {
        return jedis.zremrangeByScore(bytes,v,v2);
    }

    @Override
    public Long zremrangeByScore(byte[] bytes, byte[] bytes2, byte[] bytes3) {
        return jedis.zremrangeByScore(bytes,bytes2,bytes3);
    }

    @Override
    public Long zlexcount(byte[] bytes, byte[] bytes2, byte[] bytes3) {
        return jedis.zlexcount(bytes,bytes2,bytes3);
    }

    @Override
    public Set<byte[]> zrangeByLex(byte[] bytes, byte[] bytes2, byte[] bytes3) {
        return jedis.zrangeByLex(bytes,bytes2,bytes3);
    }

    @Override
    public Set<byte[]> zrangeByLex(byte[] bytes, byte[] bytes2, byte[] bytes3, int i, int i2) {
        return jedis.zrangeByLex(bytes,bytes2,bytes3,i,i2);
    }

    @Override
    public Long zremrangeByLex(byte[] bytes, byte[] bytes2, byte[] bytes3) {
        return jedis.zremrangeByLex(bytes,bytes2,bytes3);
    }

    @Override
    public Long linsert(byte[] bytes, BinaryClient.LIST_POSITION list_position, byte[] bytes2, byte[] bytes3) {
        return jedis.linsert(bytes,list_position,bytes2,bytes3);
    }

    @Override
    public Long lpushx(byte[] bytes, byte[]... bytes2) {
        return jedis.lpushx(bytes,bytes2);
    }

    @Override
    public Long rpushx(byte[] bytes, byte[]... bytes2) {
        return jedis.rpushx(bytes,bytes2);
    }

    @Override
    public List<byte[]> blpop(byte[] bytes) {
        return jedis.blpop(bytes);
    }

    @Override
    public List<byte[]> brpop(byte[] bytes) {
        return jedis.brpop(bytes);
    }

    @Override
    public Long del(byte[] bytes) {
        return jedis.del(bytes);
    }

    @Override
    public byte[] echo(byte[] bytes) {
        return jedis.echo(bytes);
    }

    @Override
    public Long move(byte[] bytes, int i) {
        return jedis.move(bytes,i);
    }

    @Override
    public Long bitcount(byte[] bytes) {
        return jedis.bitcount(bytes);
    }

    @Override
    public Long bitcount(byte[] bytes, long l, long l2) {
        return jedis.bitcount(bytes,l,l2);
    }

    @Override
    public Long pfadd(byte[] bytes, byte[]... bytes2) {
        return jedis.pfadd(bytes,bytes2);
    }

    @Override
    public long pfcount(byte[] bytes) {
        return jedis.pfcount(bytes);
    }

    @Override
    public Object eval(byte[] bytes, byte[] bytes2, byte[]... bytes3) {
        return jedis.eval(bytes,bytes2,bytes3);
    }

    @Override
    public Object eval(byte[] bytes, int i, byte[]... bytes2) {
        return jedis.eval(bytes,i,bytes2);
    }

    @Override
    public Object eval(byte[] bytes, List<byte[]> bytes2, List<byte[]> bytes3) {
        return jedis.eval(bytes,bytes2,bytes3);
    }

    @Override
    public Object eval(byte[] bytes) {
        return jedis.eval(bytes);
    }

    @Override
    public Object evalsha(byte[] bytes) {
        return jedis.evalsha(bytes);
    }

    @Override
    public Object evalsha(byte[] bytes, List<byte[]> bytes2, List<byte[]> bytes3) {
        return jedis.evalsha(bytes,bytes2,bytes3);
    }

    @Override
    public Object evalsha(byte[] bytes, int i, byte[]... bytes2) {
        return jedis.evalsha(bytes,i,bytes2);
    }

    @Override
    public List<Long> scriptExists(byte[]... bytes) {
        return jedis.scriptExists(bytes);
    }

    @Override
    public byte[] scriptLoad(byte[] bytes) {
        return jedis.scriptLoad(bytes);
    }

    @Override
    public String scriptFlush() {
        return jedis.scriptFlush();
    }

    @Override
    public String scriptKill() {
        return jedis.scriptKill();
    }

    @Override
    public void close()  {
        jedis.close();
    }

    @Override
    public Long del(byte[]... bytes) {
        return null;
    }

    @Override
    public List<byte[]> blpop(int i, byte[]... bytes) {
        return null;
    }

    @Override
    public List<byte[]> brpop(int i, byte[]... bytes) {
        return null;
    }

    @Override
    public List<byte[]> blpop(byte[]... bytes) {
        return null;
    }

    @Override
    public List<byte[]> brpop(byte[]... bytes) {
        return null;
    }

    @Override
    public Set<byte[]> keys(byte[] bytes) {
        return null;
    }

    @Override
    public List<byte[]> mget(byte[]... bytes) {
        return null;
    }

    @Override
    public String mset(byte[]... bytes) {
        return null;
    }

    @Override
    public Long msetnx(byte[]... bytes) {
        return null;
    }

    @Override
    public String rename(byte[] bytes, byte[] bytes2) {
        return null;
    }

    @Override
    public Long renamenx(byte[] bytes, byte[] bytes2) {
        return null;
    }

    @Override
    public byte[] rpoplpush(byte[] bytes, byte[] bytes2) {
        return new byte[0];
    }

    @Override
    public Set<byte[]> sdiff(byte[]... bytes) {
        return null;
    }

    @Override
    public Long sdiffstore(byte[] bytes, byte[]... bytes2) {
        return null;
    }

    @Override
    public Set<byte[]> sinter(byte[]... bytes) {
        return null;
    }

    @Override
    public Long sinterstore(byte[] bytes, byte[]... bytes2) {
        return null;
    }

    @Override
    public Long smove(byte[] bytes, byte[] bytes2, byte[] bytes3) {
        return null;
    }

    @Override
    public Long sort(byte[] bytes, SortingParams sortingParams, byte[] bytes2) {
        return null;
    }

    @Override
    public Long sort(byte[] bytes, byte[] bytes2) {
        return null;
    }

    @Override
    public Set<byte[]> sunion(byte[]... bytes) {
        return null;
    }

    @Override
    public Long sunionstore(byte[] bytes, byte[]... bytes2) {
        return null;
    }

    @Override
    public String watch(byte[]... bytes) {
        return null;
    }

    @Override
    public Long zinterstore(byte[] bytes, byte[]... bytes2) {
        return null;
    }

    @Override
    public Long zinterstore(byte[] bytes, ZParams zParams, byte[]... bytes2) {
        return null;
    }

    @Override
    public Long zunionstore(byte[] bytes, byte[]... bytes2) {
        return null;
    }

    @Override
    public Long zunionstore(byte[] bytes, ZParams zParams, byte[]... bytes2) {
        return null;
    }

    @Override
    public byte[] brpoplpush(byte[] bytes, byte[] bytes2, int i) {
        return new byte[0];
    }

    @Override
    public Long publish(byte[] bytes, byte[] bytes2) {
        return null;
    }

    @Override
    public void subscribe(BinaryJedisPubSub binaryJedisPubSub, byte[]... bytes) {

    }

    @Override
    public void psubscribe(BinaryJedisPubSub binaryJedisPubSub, byte[]... bytes) {

    }

    @Override
    public byte[] randomBinaryKey() {
        return new byte[0];
    }

    @Override
    public Long bitop(BitOP bitOP, byte[] bytes, byte[]... bytes2) {
        return null;
    }

    @Override
    public String pfmerge(byte[] bytes, byte[]... bytes2) {
        return null;
    }

    @Override
    public Long pfcount(byte[]... bytes) {
        return null;
    }

    @Override
    public String unwatch() {
        return null;
    }

    @Override
    public List<String> configGet(String s) {
        return jedis.configGet(s);
    }

    @Override
    public String configSet(String s, String s2) {
        return jedis.configSet(s,s2);
    }

    @Override
    public List<Slowlog> slowlogGet() {
        return jedis.slowlogGet();
    }

    @Override
    public List<Slowlog> slowlogGet(long l) {
        return jedis.slowlogGet(l);
    }

    @Override
    public Long objectRefcount(String s) {
        return jedis.objectRefcount(s);
    }

    @Override
    public String objectEncoding(String s) {
        return jedis.objectEncoding(s);
    }

    @Override
    public Long objectIdletime(String s) {
        return jedis.objectIdletime(s);
    }

    @Override
    public String clusterNodes() {
        return jedis.clusterNodes();
    }

    @Override
    public String clusterMeet(String s, int i) {
        return jedis.clusterMeet(s,i);
    }

    @Override
    public String clusterAddSlots(int... ints) {
        return jedis.clusterAddSlots(ints);
    }

    @Override
    public String clusterDelSlots(int... ints) {
        return jedis.clusterDelSlots(ints);
    }

    @Override
    public String clusterInfo() {
        return jedis.clusterInfo();
    }

    @Override
    public List<String> clusterGetKeysInSlot(int i, int i2) {
        return jedis.clusterGetKeysInSlot(i,i2);
    }

    @Override
    public String clusterSetSlotNode(int i, String s) {
        return jedis.clusterSetSlotNode(i,s);
    }

    @Override
    public String clusterSetSlotMigrating(int i, String s) {
        return jedis.clusterSetSlotMigrating(i,s);
    }

    @Override
    public String clusterSetSlotImporting(int i, String s) {
        return jedis.clusterSetSlotImporting(i,s);
    }

    @Override
    public String clusterSetSlotStable(int i) {
        return jedis.clusterSetSlotStable(i);
    }

    @Override
    public String clusterForget(String s) {
        return jedis.clusterForget(s);
    }

    @Override
    public String clusterFlushSlots() {
        return jedis.clusterFlushSlots();
    }

    @Override
    public Long clusterKeySlot(String s) {
        return jedis.clusterKeySlot(s);
    }

    @Override
    public Long clusterCountKeysInSlot(int i) {
        return jedis.clusterCountKeysInSlot(i);
    }

    @Override
    public String clusterSaveConfig() {
        return jedis.clusterSaveConfig();
    }

    @Override
    public String clusterReplicate(String s) {
        return jedis.clusterReplicate(s);
    }

    @Override
    public List<String> clusterSlaves(String s) {
        return jedis.clusterSlaves(s);
    }

    @Override
    public String clusterFailover() {
        return jedis.clusterFailover();
    }

    @Override
    public List<Object> clusterSlots() {
        return jedis.clusterSlots();
    }

    @Override
    public String clusterReset(JedisCluster.Reset reset) {
        return jedis.clusterReset(reset);
    }

    @Override
    public String set(String s, String s2) {
        return jedis.set(s,s2);
    }

    @Override
    public String set(String s, String s2, String s3, String s4, long l) {
        return jedis.set(s,s2,s3,s4,l);
    }

    @Override
    public String get(String s) {
        return jedis.get(s);
    }

    @Override
    public Boolean exists(String s) {
        return jedis.exists(s);
    }

    @Override
    public Long persist(String s) {
        return jedis.persist(s);
    }

    @Override
    public String type(String s) {
        return jedis.type(s);
    }

    @Override
    public Long expire(String s, int i) {
        return jedis.expire(s,i);
    }

    @Override
    public Long expireAt(String s, long l) {
        return jedis.expireAt(s,l);
    }

    @Override
    public Long ttl(String s) {
        return jedis.ttl(s);
    }

    @Override
    public Boolean setbit(String s, long l, boolean b) {
        return jedis.setbit(s,l,b);
    }

    @Override
    public Boolean setbit(String s, long l, String s2) {
        return jedis.setbit(s,l,s2);
    }

    @Override
    public Boolean getbit(String s, long l) {
        return jedis.getbit(s,l);
    }

    @Override
    public Long setrange(String s, long l, String s2) {
        return jedis.setrange(s,l,s2);
    }

    @Override
    public String getrange(String s, long l, long l2) {
        return jedis.getrange(s,l,l2);
    }

    @Override
    public String getSet(String s, String s2) {
        return jedis.getSet(s,s2);
    }

    @Override
    public Long setnx(String s, String s2) {
        return jedis.setnx(s,s2);
    }

    @Override
    public String setex(String s, int i, String s2) {
        return jedis.setex(s,i,s2);
    }

    @Override
    public Long decrBy(String s, long l) {
        return jedis.decrBy(s,l);
    }

    @Override
    public Long decr(String s) {
        return jedis.decr(s);
    }

    @Override
    public Long incrBy(String s, long l) {
        return jedis.incrBy(s,l);
    }

    @Override
    public Long incr(String s) {
        return jedis.incr(s);
    }

    @Override
    public Long append(String s, String s2) {
        return jedis.append(s,s2);
    }

    @Override
    public String substr(String s, int i, int i2) {
        return jedis.substr(s,i,i2);
    }

    @Override
    public Long hset(String s, String s2, String s3) {
        return jedis.hset(s,s2,s3);
    }

    @Override
    public String hget(String s, String s2) {
        return jedis.hget(s,s2);
    }

    @Override
    public Long hsetnx(String s, String s2, String s3) {
        return jedis.hsetnx(s,s2,s3);
    }

    @Override
    public String hmset(String s, Map<String, String> stringStringMap) {
        return jedis.hmset(s,stringStringMap);
    }

    @Override
    public List<String> hmget(String s, String... strings) {
        return jedis.hmget(s,strings);
    }

    @Override
    public Long hincrBy(String s, String s2, long l) {
        return jedis.hincrBy(s,s2,l);
    }

    @Override
    public Boolean hexists(String s, String s2) {
        return jedis.hexists(s,s2);
    }

    @Override
    public Long hdel(String s, String... strings) {
        return jedis.hdel(s,strings);
    }

    @Override
    public Long hlen(String s) {
        return jedis.hlen(s);
    }

    @Override
    public Set<String> hkeys(String s) {
        return jedis.hkeys(s);
    }

    @Override
    public List<String> hvals(String s) {
        return jedis.hvals(s);
    }

    @Override
    public Map<String, String> hgetAll(String s) {
        return jedis.hgetAll(s);
    }

    @Override
    public Long rpush(String s, String... strings) {
        return jedis.rpush(s,strings);
    }

    @Override
    public Long lpush(String s, String... strings) {
        return jedis.lpush(s,strings);
    }

    @Override
    public Long llen(String s) {
        return jedis.llen(s);
    }

    @Override
    public List<String> lrange(String s, long l, long l2) {
        return jedis.lrange(s,l,l2);
    }

    @Override
    public String ltrim(String s, long l, long l2) {
        return jedis.ltrim(s,l,l2);
    }

    @Override
    public String lindex(String s, long l) {
        return jedis.lindex(s,l);
    }

    @Override
    public String lset(String s, long l, String s2) {
        return jedis.lset(s,l,s2);
    }

    @Override
    public Long lrem(String s, long l, String s2) {
        return jedis.lrem(s,l,s2);
    }

    @Override
    public String lpop(String s) {
        return jedis.lpop(s);
    }

    @Override
    public String rpop(String s) {
        return jedis.rpop(s);
    }

    @Override
    public Long sadd(String s, String... strings) {
        return jedis.sadd(s,strings);
    }

    @Override
    public Set<String> smembers(String s) {
        return jedis.smembers(s);
    }

    @Override
    public Long srem(String s, String... strings) {
        return jedis.srem(s,strings);
    }

    @Override
    public String spop(String s) {
        return jedis.spop(s);
    }

    @Override
    public Long scard(String s) {
        return jedis.scard(s);
    }

    @Override
    public Boolean sismember(String s, String s2) {
        return jedis.sismember(s,s2);
    }

    @Override
    public String srandmember(String s) {
        return jedis.srandmember(s);
    }

    @Override
    public List<String> srandmember(String s, int i) {
        return jedis.srandmember(s,i);
    }

    @Override
    public Long strlen(String s) {
        return jedis.strlen(s);
    }

    @Override
    public Long zadd(String s, double v, String s2) {
        return jedis.zadd(s,v,s2);
    }

    @Override
    public Long zadd(String s, Map<String, Double> stringDoubleMap) {
        return jedis.zadd(s,stringDoubleMap);
    }

    @Override
    public Set<String> zrange(String s, long l, long l2) {
        return jedis.zrange(s,l,l2);
    }

    @Override
    public Long zrem(String s, String... strings) {
        return jedis.zrem(s,strings);
    }

    @Override
    public Double zincrby(String s, double v, String s2) {
        return jedis.zincrby(s,v,s2);
    }

    @Override
    public Long zrank(String s, String s2) {
        return jedis.zrank(s,s2);
    }

    @Override
    public Long zrevrank(String s, String s2) {
        return jedis.zrevrank(s,s2);
    }

    @Override
    public Set<String> zrevrange(String s, long l, long l2) {
        return jedis.zrevrange(s,l,l2);
    }

    @Override
    public Set<Tuple> zrangeWithScores(String s, long l, long l2) {
        return jedis.zrangeWithScores(s,l,l2);
    }

    @Override
    public Set<Tuple> zrevrangeWithScores(String s, long l, long l2) {
        return jedis.zrevrangeWithScores(s,l,l2);
    }

    @Override
    public Long zcard(String s) {
        return jedis.zcard(s);
    }

    @Override
    public Double zscore(String s, String s2) {
        return jedis.zscore(s,s2);
    }

    @Override
    public List<String> sort(String s) {
        return jedis.sort(s);
    }

    @Override
    public List<String> sort(String s, SortingParams sortingParams) {
        return jedis.sort(s,sortingParams);
    }

    @Override
    public Long zcount(String s, double v, double v2) {
        return jedis.zcount(s,v,v2);
    }

    @Override
    public Long zcount(String s, String s2, String s3) {
        return jedis.zcount(s,s2,s3);
    }

    @Override
    public Set<String> zrangeByScore(String s, double v, double v2) {
        return jedis.zrangeByScore(s,v,v2);
    }

    @Override
    public Set<String> zrangeByScore(String s, String s2, String s3) {
        return jedis.zrangeByScore(s,s2,s3);
    }

    @Override
    public Set<String> zrevrangeByScore(String s, double v, double v2) {
        return jedis.zrevrangeByScore(s,v,v2);
    }

    @Override
    public Set<String> zrangeByScore(String s, double v, double v2, int i, int i2) {
        return jedis.zrangeByScore(s,v,v2,i,i2);
    }

    @Override
    public Set<String> zrevrangeByScore(String s, String s2, String s3) {
        return jedis.zrevrangeByScore(s,s2,s3);
    }

    @Override
    public Set<String> zrangeByScore(String s, String s2, String s3, int i, int i2) {
        return jedis.zrangeByScore(s,s2,s3,i,i2);
    }

    @Override
    public Set<String> zrevrangeByScore(String s, double v, double v2, int i, int i2) {
        return jedis.zrevrangeByScore(s,v,v2,i,i2);
    }

    @Override
    public Set<Tuple> zrangeByScoreWithScores(String s, double v, double v2) {
        return jedis.zrangeByScoreWithScores(s,v,v2);
    }

    @Override
    public Set<Tuple> zrevrangeByScoreWithScores(String s, double v, double v2) {
        return jedis.zrevrangeByScoreWithScores(s,v,v2);
    }

    @Override
    public Set<Tuple> zrangeByScoreWithScores(String s, double v, double v2, int i, int i2) {
        return jedis.zrangeByScoreWithScores(s,v,v2,i,i2);
    }

    @Override
    public Set<String> zrevrangeByScore(String s, String s2, String s3, int i, int i2) {
        return jedis.zrevrangeByScore(s,s2,s3,i,i2);
    }

    @Override
    public Set<Tuple> zrangeByScoreWithScores(String s, String s2, String s3) {
        return jedis.zrangeByScoreWithScores(s,s2,s3);
    }

    @Override
    public Set<Tuple> zrevrangeByScoreWithScores(String s, String s2, String s3) {
        return jedis.zrevrangeByScoreWithScores(s,s2,s3);
    }

    @Override
    public Set<Tuple> zrangeByScoreWithScores(String s, String s2, String s3, int i, int i2) {
        return jedis.zrangeByScoreWithScores(s,s2,s3,i,i2);
    }

    @Override
    public Set<Tuple> zrevrangeByScoreWithScores(String s, double v, double v2, int i, int i2) {
        return jedis.zrevrangeByScoreWithScores(s,v,v2,i,i2);
    }

    @Override
    public Set<Tuple> zrevrangeByScoreWithScores(String s, String s2, String s3, int i, int i2) {
        return jedis.zrevrangeByScoreWithScores(s,s2,s3,i,i2);
    }

    @Override
    public Long zremrangeByRank(String s, long l, long l2) {
        return jedis.zremrangeByRank(s,l,l2);
    }

    @Override
    public Long zremrangeByScore(String s, double v, double v2) {
        return jedis.zremrangeByScore(s,v,v2);
    }

    @Override
    public Long zremrangeByScore(String s, String s2, String s3) {
        return jedis.zremrangeByScore(s,s2,s3);
    }

    @Override
    public Long zlexcount(String s, String s2, String s3) {
        return jedis.zlexcount(s,s2,s3);
    }

    @Override
    public Set<String> zrangeByLex(String s, String s2, String s3) {
        return jedis.zrangeByLex(s,s2,s3);
    }

    @Override
    public Set<String> zrangeByLex(String s, String s2, String s3, int i, int i2) {
        return jedis.zrangeByLex(s,s2,s3,i,i2);
    }

    @Override
    public Long zremrangeByLex(String s, String s2, String s3) {
        return jedis.zremrangeByLex(s,s2,s3);
    }

    @Override
    public Long linsert(String s, BinaryClient.LIST_POSITION list_position, String s2, String s3) {
        return jedis.linsert(s,list_position,s2,s3);
    }

    @Override
    public Long lpushx(String s, String... strings) {
        return jedis.lpushx(s,strings);
    }

    @Override
    public Long rpushx(String s, String... strings) {
        return jedis.rpushx(s,strings);
    }

    @Override
    public List<String> blpop(String s) {
        return jedis.blpop(s);
    }

    @Override
    public List<String> blpop(int i, String s) {
        return jedis.blpop(i,s);
    }

    @Override
    public List<String> brpop(String s) {
        return jedis.brpop(s);
    }

    @Override
    public List<String> brpop(int i, String s) {
        return jedis.brpop(i,s);
    }

    @Override
    public Long del(String s) {
        return jedis.del(s);
    }

    @Override
    public String echo(String s) {
        return jedis.echo(s);
    }

    @Override
    public Long move(String s, int i) {
        return jedis.move(s,i);
    }

    @Override
    public Long bitcount(String s) {
        return jedis.bitcount(s);
    }

    @Override
    public Long bitcount(String s, long l, long l2) {
        return jedis.bitcount(s,l,l2);
    }

    @Override
    public ScanResult<Map.Entry<String, String>> hscan(String s, int i) {
        return jedis.hscan(s,i);
    }

    @Override
    public ScanResult<String> sscan(String s, int i) {
        return jedis.sscan(s,i);
    }

    @Override
    public ScanResult<Tuple> zscan(String s, int i) {
        return jedis.zscan(s,i);
    }

    @Override
    public ScanResult<Map.Entry<String, String>> hscan(String s, String s2) {
        return jedis.hscan(s,s2);
    }

    @Override
    public ScanResult<String> sscan(String s, String s2) {
        return jedis.sscan(s,s2);
    }

    @Override
    public ScanResult<Tuple> zscan(String s, String s2) {
        return jedis.zscan(s,s2);
    }

    @Override
    public Long pfadd(String s, String... strings) {
        return jedis.pfadd(s,strings);
    }

    @Override
    public long pfcount(String s) {
        return jedis.pfcount(s);
    }

    @Override
    public Long del(String... strings) {
        return jedis.del(strings);
    }

    @Override
    public List<String> blpop(int i, String... strings) {
        return jedis.blpop(i,strings);
    }

    @Override
    public List<String> brpop(int i, String... strings) {
        return jedis.brpop(i,strings);
    }

    @Override
    public List<String> blpop(String... strings) {
        return jedis.blpop(strings);
    }

    @Override
    public List<String> brpop(String... strings) {
        return jedis.brpop(strings);
    }

    @Override
    public Set<String> keys(String s) {
        return jedis.keys(s);
    }

    @Override
    public List<String> mget(String... strings) {
        return jedis.mget(strings);
    }

    @Override
    public String mset(String... strings) {
        return jedis.mset(strings);
    }

    @Override
    public Long msetnx(String... strings) {
        return jedis.msetnx(strings);
    }

    @Override
    public String rename(String s, String s2) {
        return jedis.rename(s,s2);
    }

    @Override
    public Long renamenx(String s, String s2) {
        return jedis.renamenx(s,s2);
    }

    @Override
    public String rpoplpush(String s, String s2) {
        return jedis.rpoplpush(s,s2);
    }

    @Override
    public Set<String> sdiff(String... strings) {
        return jedis.sdiff(strings);
    }

    @Override
    public Long sdiffstore(String s, String... strings) {
        return jedis.sdiffstore(s,strings);
    }

    @Override
    public Set<String> sinter(String... strings) {
        return jedis.sinter(strings);
    }

    @Override
    public Long sinterstore(String s, String... strings) {
        return jedis.sinterstore(s,strings);
    }

    @Override
    public Long smove(String s, String s2, String s3) {
        return jedis.smove(s,s2,s3);
    }

    @Override
    public Long sort(String s, SortingParams sortingParams, String s2) {
        return jedis.sort(s,sortingParams,s2);
    }

    @Override
    public Long sort(String s, String s2) {
        return jedis.sort(s,s2);
    }

    @Override
    public Set<String> sunion(String... strings) {
        return jedis.sunion(strings);
    }

    @Override
    public Long sunionstore(String s, String... strings) {
        return jedis.sunionstore(s,strings);
    }

    @Override
    public String watch(String... strings) {
        return jedis.watch(strings);
    }

    @Override
    public Long zinterstore(String s, String... strings) {
        return jedis.zinterstore(s,strings);
    }

    @Override
    public Long zinterstore(String s, ZParams zParams, String... strings) {
        return jedis.zinterstore(s,zParams,strings);
    }

    @Override
    public Long zunionstore(String s, String... strings) {
        return jedis.zunionstore(s,strings);
    }

    @Override
    public Long zunionstore(String s, ZParams zParams, String... strings) {
        return jedis.zunionstore(s,zParams,strings);
    }

    @Override
    public String brpoplpush(String s, String s2, int i) {
        return jedis.brpoplpush(s,s2,i);
    }

    @Override
    public Long publish(String s, String s2) {
        return jedis.publish(s,s2);
    }

    @Override
    public void subscribe(JedisPubSub jedisPubSub, String... strings) {
        jedis.subscribe(jedisPubSub,strings);
    }

    @Override
    public void psubscribe(JedisPubSub jedisPubSub, String... strings) {
        jedis.psubscribe(jedisPubSub,strings);
    }

    @Override
    public String randomKey() {
        return jedis.randomKey();
    }

    @Override
    public Long bitop(BitOP bitOP, String s, String... strings) {
        return jedis.bitop(bitOP,s,strings);
    }

    @Override
    public ScanResult<String> scan(int i) {
        return jedis.scan(i);
    }

    @Override
    public ScanResult<String> scan(String s) {
        return jedis.scan(s);
    }

    @Override
    public String pfmerge(String s, String... strings) {
        return jedis.pfmerge(s,strings);
    }

    @Override
    public long pfcount(String... strings) {
        return jedis.pfcount(strings);
    }

    @Override
    public Object eval(String s, int i, String... strings) {
        return jedis.eval(s,i,strings);
    }

    @Override
    public Object eval(String s, List<String> strings, List<String> strings2) {
        return jedis.eval(s,strings,strings2);
    }

    @Override
    public Object eval(String s) {
        return jedis;
    }

    @Override
    public Object evalsha(String s) {
        return jedis.evalsha(s);
    }

    @Override
    public Object evalsha(String s, List<String> strings, List<String> strings2) {
        return jedis.evalsha(s,strings,strings2);
    }

    @Override
    public Object evalsha(String s, int i, String... strings) {
        return jedis.evalsha(s,i,strings);
    }

    @Override
    public Boolean scriptExists(String s) {
        return jedis.scriptExists(s);
    }

    @Override
    public List<Boolean> scriptExists(String... strings) {
        return jedis.scriptExists(strings);
    }

    @Override
    public String scriptLoad(String s) {
        return jedis.scriptLoad(s);
    }

    @Override
    public List<Map<String, String>> sentinelMasters() {
        return jedis.sentinelMasters();
    }

    @Override
    public List<String> sentinelGetMasterAddrByName(String s) {
        return jedis.sentinelGetMasterAddrByName(s);
    }

    @Override
    public Long sentinelReset(String s) {
        return jedis.sentinelReset(s);
    }

    @Override
    public List<Map<String, String>> sentinelSlaves(String s) {
        return jedis.sentinelSlaves(s);
    }

    @Override
    public String sentinelFailover(String s) {
        return jedis.sentinelFailover(s);
    }

    @Override
    public String sentinelMonitor(String s, String s2, int i, int i2) {
        return jedis.sentinelMonitor(s,s2,i,i2);
    }

    @Override
    public String sentinelRemove(String s) {
        return jedis.sentinelRemove(s);
    }

    @Override
    public String sentinelSet(String s, Map<String, String> stringStringMap) {
        return jedis.sentinelSet(s,stringStringMap);
    }


}
