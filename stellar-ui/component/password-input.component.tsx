import { Button, Input, InputGroup, InputRightElement } from "@chakra-ui/react";
import React from "react";

interface PasswordInputProps {
  placeHolderText: string;
}

const PasswordInput: React.FC<PasswordInputProps> = ({ placeHolderText }) => {
  const [show, setShow] = React.useState(false);
  const handleClick = () => setShow(!show);

  return (
    <InputGroup size="md">
      <Input
        pr="4.5rem"
        type={show ? "text" : "password"}
        placeholder={placeHolderText}
      />
      <InputRightElement width="4.5rem">
        <Button h="1.75rem" size="sm" is onClick={handleClick}>
          {show ? "Hide" : "Show"}
        </Button>
      </InputRightElement>
    </InputGroup>
  );
};

export default PasswordInput;
