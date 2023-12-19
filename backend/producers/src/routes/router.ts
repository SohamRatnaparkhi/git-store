import { Router } from "express";
import webHooksRouter from "./webhooks/webhooks";

const router = Router();

router.use("/webhooks", webHooksRouter);

export default router;