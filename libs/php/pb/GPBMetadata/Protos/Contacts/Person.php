<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: protos/contacts/person.proto

namespace GPBMetadata\Protos\Contacts;

class Person
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
�
protos/contacts/person.protoContacts"�
Person
name (	

id (
email (	,
phones (2.Contacts.Person.PhoneNumber0
last_updated (2.google.protobuf.TimestampG
PhoneNumber
number (	(
type (2.Contacts.Person.PhoneType"+
	PhoneType

MOBILE 
HOME
WORKB)Zpb/contacts�Pb.Contacts�Pb/Contactsbproto3'
        , true);

        static::$is_initialized = true;
    }
}

