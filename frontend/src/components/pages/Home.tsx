import React, { useEffect } from 'react'
import { Flex, Heading } from '@chakra-ui/react'
import { useDispatch } from 'react-redux'
import Auth from '../atoms/Auth'
import ActionButton from '../atoms/ActionButton'
import { setOpenStatus } from '../../reducks/modals/slices'
import AddToDoModal from '../organisms/AddToDoModal'

const Home = () => {
    const dispatch = useDispatch()
    const action = () => {
        dispatch(setOpenStatus({ isOpen: true }))
    }

    return (
        <Auth>
            <Flex w="100%" h="100vh" display="flex" align="center" justify="center" pt="56px">
                <Heading as="h1" size="2xl">
                    Index Page.
                </Heading>
                <ActionButton action={action} />
                <AddToDoModal />
            </Flex>
        </Auth>
    )
}

export default Home
