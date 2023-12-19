import { Router } from "express";
import githubRouter from "./github";

const router = Router();

router.use("/github", githubRouter);

export default router; 