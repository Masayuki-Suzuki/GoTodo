import React, { useEffect, useRef } from 'react'
import { useNavigate } from 'react-router-dom'
import axios from 'axios'
import { VerifyUser } from '../../reducks/users/types'

type AuthPropsType = {
    children: JSX.Element
}

const Auth = ({ children }: AuthPropsType) => {
    const navigate = useNavigate()
    const isFetched = useRef(false)

    const checkAuth = async () => {
        const token = sessionStorage.getItem('token')
        if (!token) {
            return false
        }
        try {
            const { data } = await axios.post<VerifyUser>('http://localhost:4000/api/admin/verify-user', { token })
            return data.is_authorised
        } catch (e) {
            console.error(e)
            return false
        }
    }

    useEffect(() => {
        if (!isFetched.current) {
            isFetched.current = true
            const isAuthorised = checkAuth()

            if (!isAuthorised) {
                navigate('/login')
            }
        }
    }, [])

    return children
}

export default Auth
