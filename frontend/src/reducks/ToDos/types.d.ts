import { Nullable } from '~/types/utils'

export type ToDo = {
    title: string
    description: string
    dueDate: string
    status: string
    createdAt: string
    completedAt: string
    userID: number
}

export type ToDos = {
    todos: ToDo[]
}

export type ToDoField = {
    title: string
    description: string
    dueDate: Nullable<string>
}
