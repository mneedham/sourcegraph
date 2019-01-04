import { map } from 'rxjs/operators'
import { ViewComponent, Window } from 'sourcegraph'
import { MessageType } from '../client/services/notifications'
import { Position } from '../extension/types/position'
import { Selection } from '../extension/types/selection'
import { assertToJSON } from '../extension/types/testHelpers'
import { collectSubscribableValues, integrationTestContext } from './testHelpers'

describe('Windows (integration)', () => {
    describe('app.activeWindow', () => {
        test('returns the active window', async () => {
            const { extensionHost } = await integrationTestContext()
            const viewComponent: Pick<ViewComponent, 'type' | 'document'> = {
                type: 'CodeEditor' as 'CodeEditor',
                document: { uri: 'file:///f', languageId: 'l', text: 't' },
            }
            assertToJSON(extensionHost.app.activeWindow, {
                visibleViewComponents: [viewComponent],
                activeViewComponent: viewComponent,
            } as Window)
        })
    })

    describe('app.windows', () => {
        test('lists windows', async () => {
            const { extensionHost } = await integrationTestContext()
            const viewComponent: Pick<ViewComponent, 'type' | 'document'> = {
                type: 'CodeEditor' as 'CodeEditor',
                document: { uri: 'file:///f', languageId: 'l', text: 't' },
            }
            assertToJSON(extensionHost.app.windows, [
                {
                    visibleViewComponents: [viewComponent],
                    activeViewComponent: viewComponent,
                },
            ] as Window[])
        })

        test('adds new text documents', async () => {
            const { model, extensionHost } = await integrationTestContext()

            model.next({
                ...model.value,
                visibleViewComponents: [
                    {
                        type: 'textEditor',
                        item: { uri: 'file:///f2', languageId: 'l2', text: 't2' },
                        selections: [],
                        isActive: true,
                    },
                ],
            })
            await extensionHost.internal.sync()

            const viewComponent: Pick<ViewComponent, 'type' | 'document'> = {
                type: 'CodeEditor' as 'CodeEditor',
                document: { uri: 'file:///f2', languageId: 'l2', text: 't2' },
            }
            assertToJSON(extensionHost.app.windows, [
                {
                    visibleViewComponents: [viewComponent],
                    activeViewComponent: viewComponent,
                },
            ] as Window[])
        })
    })

    describe('Window', () => {
        test('Window#visibleViewComponent', async () => {
            const { model, extensionHost } = await integrationTestContext()

            model.next({
                ...model.value,
                visibleViewComponents: [
                    {
                        type: 'textEditor',
                        item: {
                            uri: 'file:///inactive',
                            languageId: 'inactive',
                            text: 'inactive',
                        },
                        selections: [],
                        isActive: false,
                    },
                    ...(model.value.visibleViewComponents || []),
                ],
            })
            await extensionHost.internal.sync()

            assertToJSON(extensionHost.app.windows[0].visibleViewComponents, [
                {
                    type: 'CodeEditor' as 'CodeEditor',
                    document: { uri: 'file:///inactive', languageId: 'inactive', text: 'inactive' },
                },
                {
                    type: 'CodeEditor' as 'CodeEditor',
                    document: { uri: 'file:///f', languageId: 'l', text: 't' },
                },
            ] as ViewComponent[])
        })

        test('Window#activeViewComponent', async () => {
            const { model, extensionHost } = await integrationTestContext()

            model.next({
                ...model.value,
                visibleViewComponents: [
                    {
                        type: 'textEditor',
                        item: {
                            uri: 'file:///inactive',
                            languageId: 'inactive',
                            text: 'inactive',
                        },
                        selections: [],
                        isActive: false,
                    },
                    ...(model.value.visibleViewComponents || []),
                ],
            })
            await extensionHost.internal.sync()

            assertToJSON(extensionHost.app.windows[0].activeViewComponent, {
                type: 'CodeEditor' as 'CodeEditor',
                document: { uri: 'file:///f', languageId: 'l', text: 't' },
            } as ViewComponent)
        })

        test('Window#subscribe', async () => {
            const { model, extensionHost } = await integrationTestContext()

            const values: Window[] = []
            extensionHost.app.windows[0].subscribe(win => values.push(win))

            model.next({
                ...model.value,
                visibleViewComponents: [
                    {
                        type: 'textEditor',
                        item: {
                            uri: 'file:///x',
                            languageId: '',
                            text: '',
                        },
                        selections: [new Selection(new Position(1, 2), new Position(3, 4))],
                        isActive: true,
                    },
                    ...(model.value.visibleViewComponents || []),
                ],
            })
            await extensionHost.internal.sync()

            model.next({
                ...model.value,
                visibleViewComponents: [
                    {
                        type: 'textEditor',
                        item: {
                            uri: 'file:///x',
                            languageId: 'l',
                            text: 't',
                        },
                        selections: [],
                        isActive: true,
                    },
                    ...(model.value.visibleViewComponents || []),
                ],
            })
            await extensionHost.internal.sync()

            const want: {
                activeViewComponent: Pick<ViewComponent, 'type' | 'document' | 'selections'> | undefined
            }[] = [
                {
                    activeViewComponent: {
                        type: 'CodeEditor' as 'CodeEditor',
                        document: { uri: 'file:///f', languageId: 'l', text: 't' },
                        selections: [],
                    },
                },
                {
                    activeViewComponent: {
                        type: 'CodeEditor' as 'CodeEditor',
                        document: { uri: 'file:///f', languageId: 'l', text: 't' },
                        selections: [new Selection(new Position(1, 2), new Position(3, 4))],
                    },
                },
            ]
            assertToJSON(values, want)
        })

        test('Window#showNotification', async () => {
            const { extensionHost, services } = await integrationTestContext()
            const values = collectSubscribableValues(services.notifications.showMessages)
            extensionHost.app.activeWindow!.showNotification('a') // tslint:disable-line deprecation
            await extensionHost.internal.sync()
            expect(values).toEqual([{ message: 'a', type: MessageType.Info }] as typeof values)
        })

        test('Window#showMessage', async () => {
            const { extensionHost, services } = await integrationTestContext()
            services.notifications.showMessageRequests.subscribe(({ resolve }) => resolve(Promise.resolve(null)))
            const values = collectSubscribableValues(
                services.notifications.showMessageRequests.pipe(map(({ message, type }) => ({ message, type })))
            )
            expect(await extensionHost.app.activeWindow!.showMessage('a')).toBe(null)
            expect(values).toEqual([{ message: 'a', type: MessageType.Info }] as typeof values)
        })

        test('Window#showInputBox', async () => {
            const { extensionHost, services } = await integrationTestContext()
            services.notifications.showInputs.subscribe(({ resolve }) => resolve(Promise.resolve('c')))
            const values = collectSubscribableValues(
                services.notifications.showInputs.pipe(map(({ message, defaultValue }) => ({ message, defaultValue })))
            )
            expect(await extensionHost.app.activeWindow!.showInputBox({ prompt: 'a', value: 'b' })).toBe('c')
            expect(values).toEqual([{ message: 'a', defaultValue: 'b' }] as typeof values)
        })
    })
})
