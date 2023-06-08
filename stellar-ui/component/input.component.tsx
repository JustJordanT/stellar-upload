import React from "react";
import { Input } from "@chakra-ui/react";

interface InputProps {
  inputText: string;
}

const InputComponent: React.FC<InputProps> = ({ inputText }) => {
  return <Input placeholder={inputText} />;
};

export default InputComponent;
