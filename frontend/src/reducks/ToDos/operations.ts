import axios from 'axios'
import { Dispatch } from 'react'
import { Action } from '@reduxjs/toolkit'
import { loadingStatusAction, setFetchError } from '../fetchStatus/slices'
import { setOpenStatus } from '../modals/slices'
import { ToDo, ToDoField } from './types'
import { addToDo as addToDoToState, updateAllToDo } from './slices'

export const addToDo = (fieldData: ToDoField) => {
    return async (dispatch: Dispatch<Action>) => {
        try {
            dispatch(loadingStatusAction({ isLoading: true }))
            const postData = {
                ...fieldData,
                token: sessionStorage.getItem('token')
            }
            const { data } = await axios.post('/todo/create', postData)
            const todo: ToDo = data
            dispatch(addToDoToState(todo))
            console.log(data)
            dispatch(loadingStatusAction({ isLoading: false }))
            dispatch(setOpenStatus({ isOpen: false }))
        } catch (e) {
            if (axios.isAxiosError(e)) {
                if (e.response) {
                    dispatch(setFetchError({ error: e.response.data.error.message }))
                }
            }
            dispatch(loadingStatusAction({ isLoading: false }))
        }
    }
}

export const getAllToDo = () => {
    return async (dispatch: Dispatch<Action>) => {
        const token = sessionStorage.getItem('token')
        try {
            dispatch(loadingStatusAction({ isLoading: true }))
            const { data } = await axios.get('/todo/all', {
                headers: {
                    Authorization: token
                }
            })
            dispatch(updateAllToDo(data))
            dispatch(loadingStatusAction({ isLoading: false }))
        } catch (e) {
            if (axios.isAxiosError(e)) {
                if (e.response) {
                    dispatch(setFetchError({ error: e.response.data.error.message }))
                }
            }
            dispatch(loadingStatusAction({ isLoading: false }))
        }
    }
}
