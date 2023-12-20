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
                clientId: 'Iv1.bb50ab8e334c6d8f',
                clientSecret: '361b8361c480fc03a0081b832434dad2c3c23372'
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