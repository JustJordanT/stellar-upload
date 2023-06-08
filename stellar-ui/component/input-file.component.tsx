import React from "react";
import { Input } from "@chakra-ui/react";

interface InputProps {
  inputText: string;
}

const InputFileComponent: React.FC<InputProps> = ({ inputText }) => {
  return <Input placeholder={inputText} type="file" />;
};

export default InputFileComponent;
