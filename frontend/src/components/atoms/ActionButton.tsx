import React from 'react'
import { Flex } from '@chakra-ui/react'
import { AddIcon } from '@chakra-ui/icons'

type ActionButtonProps = {
    action?: () => void
}

const ActionButton = ({ action }: ActionButtonProps) => {
    return (
        <Flex
            onClick={action}
            borderRadius="100%"
            bg="#19448e"
            w={12}
            h={12}
            cursor="pointer"
            boxShadow="lg"
            pos="absolute"
            bottom={6}
            right={6}
            align="center"
            justify="center"
            transition="0.2s"
            _hover={{ transform: 'scale(1.3) rotate(180deg)' }}
        >
            <AddIcon color="white" fontSize={16} />
        </Flex>
    )
}

export default ActionButton
