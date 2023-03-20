import axios from 'axios'
import { Action } from '@reduxjs/toolkit'
import { Dispatch } from 'react'
import { loadingStatusAction, setFetchError } from '../fetchStatus/slices'
import { LogInField, ResponseUserData, SignUpField, User } from '../users/types'
import { signInAction } from './slices'

const setUpUserData = (dispatch: Dispatch<Action>, data: ResponseUserData) => {
    const userData: User = {
        id: data.user.id,
        firstName: data.user.first_name,
        lastName: data.user.last_name,
        fullName: data.user.full_name,
        emailAddress: data.user.email
    }
    sessionStorage.setItem('token', data.user.token)
    dispatch(signInAction(userData))
    dispatch(loadingStatusAction({ isLoading: false }))
    window.location.href = '/'
}

export const signUp = (fieldData: SignUpField) => {
    return async (dispatch: Dispatch<Action>) => {
        try {
            dispatch(loadingStatusAction({ isLoading: true }))
            const { data } = await axios.post<ResponseUserData>('/admin/register', fieldData)
            setUpUserData(dispatch, data)
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

export const login = (fieldData: LogInField) => {
    return async (dispatch: Dispatch<Action>) => {
        try {
            dispatch(loadingStatusAction({ isLoading: true }))
            const { data } = await axios.post('/admin/login', fieldData)
            setUpUserData(dispatch, data)
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

export default {
    signUp
}
