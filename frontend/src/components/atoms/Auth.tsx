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
            navigate('/login')
        } else {
            try {
                await axios.post<VerifyUser>('http://localhost:4000/api/admin/verify-user', {
                    token
                })
            } catch (e) {
                sessionStorage.removeItem('token')
                navigate('/login')
            }
        }
    }

    useEffect(() => {
        if (!isFetched.current) {
            isFetched.current = true
            void checkAuth()
        }
    }, [])

    return children
}

export default Auth
