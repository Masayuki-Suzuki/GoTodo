import React from 'react'
import {
    Button,
    Modal,
    ModalBody,
    ModalCloseButton,
    ModalContent,
    ModalFooter,
    ModalHeader,
    ModalOverlay
} from '@chakra-ui/react'
import { useDispatch, useSelector } from 'react-redux'
import { getModalStatus } from '../../reducks/modals/selectors'
import { setOpenStatus } from '../../reducks/modals/slices'

const AddToDoModal = () => {
    const isOpen = useSelector(getModalStatus)
    const dispatch = useDispatch()
    const onCloseHandler = () => {
        dispatch(setOpenStatus({ isOpen: false }))
    }

    return (
        <Modal isOpen={isOpen} onClose={onCloseHandler}>
            <ModalOverlay />
            <ModalContent>
                <ModalHeader textAlign="center">Add New ToDo</ModalHeader>
                <ModalCloseButton />
                <ModalBody>モーダル！</ModalBody>
                <ModalFooter>
                    <Button mr={3} variant="ghost" onClick={onCloseHandler}>
                        Cancel
                    </Button>
                    <Button colorScheme="blue" onClick={onCloseHandler}>
                        Add
                    </Button>
                </ModalFooter>
            </ModalContent>
        </Modal>
    )
}

export default AddToDoModal
