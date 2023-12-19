import { Request } from "express"
import { RequestStatus } from "../../constants/status";
import { filterPayload } from "../../helpers/github/filter";
import { produceToKafkaTopic } from "../../helpers/kafka/producer";
import { ServerResponse } from "src/types/server";

export const produceGithubWebhookPayload = async (req: Request, res: ServerResponse<any>) => {
    const filteredPayload = filterPayload(req);
    const reqStatus = new RequestStatus();
    if (filteredPayload.type === "pull_request") {
        const kafkaResponse = await produceToKafkaTopic("github-pull-requests", filteredPayload);
        return res.status(reqStatus.OK.code).json({
            data: kafkaResponse,
            statusType: reqStatus.OK,
        });
    }
    if (filteredPayload.type === "issues") {
        const kafkaResponse = await produceToKafkaTopic("github-issues", filteredPayload);
        return res.status(reqStatus.OK.code).json({
            data: kafkaResponse,
            statusType: reqStatus.OK,
        });
    }
    return res.status(reqStatus.OK.code).json({
        data: "No payload",
        statusType: reqStatus.OK,
    });
}