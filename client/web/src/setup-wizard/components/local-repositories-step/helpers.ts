import { open } from '@tauri-apps/api/dialog'

import { GetLocalCodeHostsResult } from '../../../graphql-operations'

export interface LocalCodeHost {
    id: string
    path: string
    autogenerated: boolean
}

/**
 * Parses gql response and returns paths of all non-autogenerated local services.
 */
export function getLocalServicePaths(data?: GetLocalCodeHostsResult): string[] {
    const localCodeHosts = getLocalServices(data)

    if (!localCodeHosts) {
        return []
    }

    return localCodeHosts.map(item => item.path)
}

/**
 * Returns the local services that have been created manually by user in the setup wizard,
 * it ignores autogenerated on the backend local external services
 */
export function getLocalServices(data?: GetLocalCodeHostsResult, isAutogenerated?: boolean): LocalCodeHost[] {
    if (!data) {
        return []
    }

    return (
        data.localExternalServices.filter(service =>
            isAutogenerated ? service.autogenerated : !service.autogenerated
        ) ?? []
    )
}

/**
 * Generates minimal code host configuration for the local repositories code host.
 */
export function createDefaultLocalServiceConfig(path: string): string {
    return `{ "url":"${window.context.srcServeGitUrl}", "root": "${path}", "repos": ["src-serve-local"] }`
}

type Path = string

/**
 * Calls native file picker window, returns list of picked files paths.
 * In case if picker was closed/canceled returns null
 */
export async function callFilePicker(): Promise<Path[] | null> {
    const selected = await open({
        directory: true,
        multiple: true,
    })

    if (Array.isArray(selected)) {
        return selected
    }

    if (selected !== null) {
        return [selected]
    }

    return null
}
