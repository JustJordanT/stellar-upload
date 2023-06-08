import React from "react";
import { Button, ButtonGroup } from "@chakra-ui/react";

interface buttonProps {
  buttonText: string;
}

const ButtonComponent: React.FC<buttonProps> = ({ buttonText }) => {
  return (
    <div>
      <Button colorScheme="blue">{buttonText}</Button>
    </div>
  );
};

export default ButtonComponent;
