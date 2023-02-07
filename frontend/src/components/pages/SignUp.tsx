import React, { useEffect } from 'react'
import { useNavigate } from 'react-router-dom'
import { Flex } from '@chakra-ui/react'
import SignUpCard from '../organisms/SignUpCard'

const SignUpPage = () => {
    const navigate = useNavigate()

    useEffect(() => {
        const token = sessionStorage.getItem('token')
        if (token) {
            navigate('/')
        }
    }, [])

    return (
        <Flex w="100%" h="100vh" display="flex" align="center" justify="center" pt="56px">
            <SignUpCard />
        </Flex>
    )
}

export default SignUpPage
