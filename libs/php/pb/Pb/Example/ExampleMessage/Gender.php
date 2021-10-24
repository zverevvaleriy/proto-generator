<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: protos/example/example.proto

namespace Pb\Example\ExampleMessage;

use UnexpectedValueException;

/**
 * Protobuf type <code>Example.ExampleMessage.Gender</code>
 */
class Gender
{
    /**
     * Generated from protobuf enum <code>NONE = 0;</code>
     */
    const NONE = 0;
    /**
     * Generated from protobuf enum <code>MALE = 1;</code>
     */
    const MALE = 1;
    /**
     * Generated from protobuf enum <code>FEMALE = 2;</code>
     */
    const FEMALE = 2;

    private static $valueToName = [
        self::NONE => 'NONE',
        self::MALE => 'MALE',
        self::FEMALE => 'FEMALE',
    ];

    public static function name($value)
    {
        if (!isset(self::$valueToName[$value])) {
            throw new UnexpectedValueException(sprintf(
                    'Enum %s has no name defined for value %s', __CLASS__, $value));
        }
        return self::$valueToName[$value];
    }


    public static function value($name)
    {
        $const = __CLASS__ . '::' . strtoupper($name);
        if (!defined($const)) {
            throw new UnexpectedValueException(sprintf(
                    'Enum %s has no value defined for name %s', __CLASS__, $name));
        }
        return constant($const);
    }
}

// Adding a class alias for backwards compatibility with the previous class name.
class_alias(Gender::class, \Pb\Example\ExampleMessage_Gender::class);

