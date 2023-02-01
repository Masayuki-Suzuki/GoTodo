import axios from 'axios'
import { Action } from '@reduxjs/toolkit'
import { push } from 'connected-react-router'
import { Dispatch } from 'react'
import { signInAction } from './slices'
import { SignUpField } from '~/reducks/users/types'
import { loadingStatusAction } from '../fetchStatus/slices'

export const signUp = (fieldData: SignUpField) => {
    return async (dispatch: Dispatch<Action>) => {
        try {
            dispatch(
                loadingStatusAction({
                    isLoading: true
                })
            )
            const res = await axios.post('http://localhost:4000/api/admin/register', fieldData)
            console.log(res)
            dispatch(
                loadingStatusAction({
                    isLoading: false
                })
            )
            dispatch(push('/'))
        } catch (e) {
            console.error(e)
            dispatch(
                loadingStatusAction({
                    isLoading: false
                })
            )
        }
    }
}
