import express from 'express';
import { Send } from 'express-serve-static-core';

interface Body<T> {
    data: !T
    error: !any
    StatusType: StatusType
}

interface ServerResponse<T> extends express.Response {
    json: Send<Body<T>, this>
}