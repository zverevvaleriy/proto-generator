<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: protos/contacts/address_book.proto

namespace GPBMetadata\Protos\Contacts;

class AddressBook
{
    public static $is_initialized = false;

    public static function initOnce() {
        $pool = \Google\Protobuf\Internal\DescriptorPool::getGeneratedPool();

        if (static::$is_initialized == true) {
          return;
        }
        \GPBMetadata\Protos\Contacts\Person::initOnce();
        $pool->internalAddGeneratedFile(
            '
�
"protos/contacts/address_book.protoContacts"/
AddressBook 
people (2.Contacts.PersonBLZ<github.com/zverevvaleriy/proto-generator/libs/go/pb/contacts�Pb.Contactsbproto3'
        , true);

        static::$is_initialized = true;
    }
}
