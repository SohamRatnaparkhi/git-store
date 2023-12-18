import express from 'express';
import { Send } from 'express-serve-static-core';
import { StatusType } from './global';

interface Body<T> {
    data?: T
    error?: any
    StatusType: StatusType
}

export interface ServerResponse<T> extends express.Response {
    json: Send<Body<T>, this>
}