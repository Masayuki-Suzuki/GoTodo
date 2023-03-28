import { createSlice, PayloadAction } from '@reduxjs/toolkit'
import { ToDos, ToDo, ToDoPayload } from './types'

const initialState: ToDos = {
    todos: []
}

export const todoSlice = createSlice({
    name: 'todo',
    initialState,
    reducers: {
        addToDo: (state: ToDos, action: PayloadAction<ToDo>) => {
            const updatedToDoAry = [...state.todos, action.payload]
            const updatedData: ToDos = { todos: updatedToDoAry }
            return updatedData
        },
        updateAllToDo: (state: ToDos, action: PayloadAction<ToDos>) => {
            const updatedData = { ...state, ...action.payload }
            return updatedData
        }
    }
})

export const { addToDo, updateAllToDo } = todoSlice.actions
export const todoState = todoSlice.reducer
