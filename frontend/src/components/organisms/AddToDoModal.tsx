import React from 'react'
import { Modal, ModalCloseButton, ModalContent, ModalHeader, ModalOverlay } from '@chakra-ui/react'
import { useDispatch, useSelector } from 'react-redux'
import { getModalStatus } from '../../reducks/modals/selectors'
import { setOpenStatus } from '../../reducks/modals/slices'
import AddToDoForm from '../organisms/AddToDoForm'

const AddToDoModal = () => {
    const isOpen = useSelector(getModalStatus)
    const dispatch = useDispatch()
    const onCloseHandler = () => {
        dispatch(setOpenStatus({ isOpen: false }))
    }

    return (
        <Modal isOpen={isOpen} onClose={onCloseHandler} size="xl" isCentered>
            <ModalOverlay />
            <ModalContent>
                <ModalHeader textAlign="center">Add New ToDo</ModalHeader>
                <ModalCloseButton />
                <AddToDoForm />
            </ModalContent>
        </Modal>
    )
}

export default AddToDoModal
