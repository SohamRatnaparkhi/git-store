import { App } from 'octokit'

export class GithubApp {
    private static instance: GithubApp;
    private _app = {} as App;
    private constructor() {
        const ghApp = new App({
            appId: process.env.GH_APP_ID || '',
            privateKey: process.env.GH_APP_PRIVATE_KEY || '',
            webhooks: {
                secret: process.env.GH_APP_WEBHOOK_SECRET || ''
            },
            oauth: {
                clientId: process.env.GH_APP_CLIENT_ID || '', 
                clientSecret: process.env.GH_APP_CLIENT_SECRET || ''
            }
        })
        this._app = ghApp;
    }
    public static getInstance() : GithubApp {
        if (!GithubApp.instance) {
            const ghApp = new GithubApp();
            GithubApp.instance = ghApp;
        }
        return GithubApp.instance;
    }
    get app() : App {
        return this._app
    }
}