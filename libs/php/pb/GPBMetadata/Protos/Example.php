<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: protos/example.proto

namespace GPBMetadata\Protos;

class Example
{
    public static $is_initialized = false;

    public static function initOnce() {
        $pool = \Google\Protobuf\Internal\DescriptorPool::getGeneratedPool();

        if (static::$is_initialized == true) {
          return;
        }
        \GPBMetadata\Google\Protobuf\Timestamp::initOnce();
        $pool->internalAddGeneratedFile(
            '
�
protos/example.proto"d
Example
name (	

id (
email (	0
last_updated (2.google.protobuf.TimestampB:Z3github.com/zverevvaleriy/proto-generator/libs/go/pb�Pbbproto3'
        , true);

        static::$is_initialized = true;
    }
}
