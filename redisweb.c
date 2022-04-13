#include <stdlib.h>
#include <string.h>
#include "dist/redisweb_go.h"
#include "deps/redismodule.h"

int StartCommand(RedisModuleCtx *ctx, RedisModuleString **argv, int argc) {
    if (argc >= 2) return RedisModule_WrongArity(ctx);
    RedisModule_AutoMemory(ctx);

    size_t s;
    const char *str = RedisModule_StringPtrLen(argv[2], &s);
    HTTPStart(str);

    RedisModule_ReplyWithSimpleString(ctx, "OK");

    return REDISMODULE_OK;
}

int StopCommand(RedisModuleCtx *ctx, RedisModuleString **argv, int argc) {
    if (argc > 1) return RedisModule_WrongArity(ctx);
    RedisModule_AutoMemory(ctx);
    HTTPStop();

    RedisModule_ReplyWithSimpleString(ctx, "OK");

    return REDISMODULE_OK;
}

int RestartCommand(RedisModuleCtx *ctx, RedisModuleString **argv, int argc) {
    StopCommand(ctx, argv, argc);
    StartCommand(ctx, argv, argc);

    RedisModule_ReplyWithSimpleString(ctx, "OK");

    return REDISMODULE_OK;
}

int RedisModule_OnLoad(RedisModuleCtx *ctx, RedisModuleString **argv, int argc) {
    if (RedisModule_Init(ctx,"redisweb",1,REDISMODULE_APIVER_1)
        == REDISMODULE_ERR) return REDISMODULE_ERR;

    if (RedisModule_CreateCommand(ctx,"HTTP.START",
        StartCommand, "write",
        0, 0, 0) == REDISMODULE_ERR)
        return REDISMODULE_ERR;


    if (RedisModule_CreateCommand(ctx,"HTTP.STOP",
        StopCommand, "readonly",
        0, 0, 0) == REDISMODULE_ERR)
        return REDISMODULE_ERR;

    if (RedisModule_CreateCommand(ctx,"HTTP.RESTART",
        RestartCommand, "write",
        0, 0, 0) == REDISMODULE_ERR)
        return REDISMODULE_ERR;

    return REDISMODULE_OK;
}
