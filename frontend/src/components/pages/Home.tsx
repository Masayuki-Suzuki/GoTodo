import React, { useEffect } from 'react'
import { Flex, Heading } from '@chakra-ui/react'
import { useNavigate } from 'react-router-dom'
import Auth from '../atoms/Auth'

const Home = () => {
    const navigate = useNavigate()

    useEffect(() => {
        const token = sessionStorage.getItem('token')

        if (!token) {
            navigate('/login')
        }
    }, [])

    return (
        <Auth>
            <Flex w="100%" h="100vh" display="flex" align="center" justify="center" pt="56px">
                <Heading as="h1" size="2xl">
                    Index Page.
                </Heading>
            </Flex>
        </Auth>
    )
}

export default Home
