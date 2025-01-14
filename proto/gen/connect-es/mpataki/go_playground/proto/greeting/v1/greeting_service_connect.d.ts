// @generated by protoc-gen-connect-es v0.13.0 with parameter "target=js+dts"
// @generated from file mpataki/go_playground/proto/greeting/v1/greeting_service.proto (package mpataki.go_playground.proto.greeting.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { HelloRequest, HelloResponse } from "./greeting_service_pb.js";
import { MethodKind } from "@bufbuild/protobuf";

/**
 * GreetingService provides greeting functionality
 *
 * @generated from service mpataki.go_playground.proto.greeting.v1.GreetingService
 */
export declare const GreetingService: {
  readonly typeName: "mpataki.go_playground.proto.greeting.v1.GreetingService",
  readonly methods: {
    /**
     * SayHello sends a single greeting
     *
     * @generated from rpc mpataki.go_playground.proto.greeting.v1.GreetingService.SayHello
     */
    readonly sayHello: {
      readonly name: "SayHello",
      readonly I: typeof HelloRequest,
      readonly O: typeof HelloResponse,
      readonly kind: MethodKind.Unary,
    },
  }
};
