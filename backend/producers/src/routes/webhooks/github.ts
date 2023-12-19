import { Router } from "express";

import { produceGithubWebhookPayload } from "../../controllers/webhook/github";

const router = Router();

router.post("/", produceGithubWebhookPayload);
router.get("/", (_req, res) => {
    res.send("Hello world");
});

export default router;