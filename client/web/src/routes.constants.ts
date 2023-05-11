export enum PageRoutes {
    Index = '/',
    Search = '/search',
    SearchConsole = '/search/console',
    SignIn = '/sign-in',
    SignUp = '/sign-up',
    UnlockAccount = '/unlock-account/:token',
    Welcome = '/welcome',
    Settings = '/settings',
    User = '/user/*',
    Organizations = '/organizations/*',
    SiteAdmin = '/site-admin/*',
    SiteAdminInit = '/site-admin/init',
    PasswordReset = '/password-reset',
    ApiConsole = '/api/console',
    UserArea = '/users/:username/*',
    Survey = '/survey/:score?',
    Extensions = '/extensions',
    Help = '/help/*',
    Debug = '/-/debug/*',
    RepoContainer = '/*',
    SetupWizard = '/setup',
    InstallGitHubAppSuccess = '/install-github-app-success',
    Teams = '/teams/*',
    RequestAccess = '/request-access/*',
}

export enum EnterprisePageRoutes {
    BatchChanges = '/batch-changes/*',
    CodeMonitoring = '/code-monitoring/*',
    Insights = '/insights/*',
    Contexts = '/contexts',
    CreateContext = '/contexts/new',
    EditContext = '/contexts/:specOrOrg/:spec?/edit',
    Context = '/contexts/:specOrOrg/:spec?',
    NotebookCreate = '/notebooks/new',
    Notebook = '/notebooks/:id',
    Notebooks = '/notebooks',
    SearchNotebook = '/search/notebook',
    CodySearch = '/search/cody',
    Cody = '/cody',
    Own = '/own',
    AppComingSoon = '/app/coming-soon',
    AppAuthCallback = '/app/auth/callback',
    Sentinel = '/sentinel/*',
}
