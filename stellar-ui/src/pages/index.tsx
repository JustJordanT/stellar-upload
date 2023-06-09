import Image from "next/image";
import { Inter } from "next/font/google";
import { ChakraProvider, Container } from "@chakra-ui/react";
import SiteTitle from "../../component/title.component";
import PasswordInput from "../../component/password-input.component";
import InputComponent from "../../component/input.component";
import InputFileComponent from "../../component/input-file.component";
import ButtonComponent from "../../component/button.component";

const inter = Inter({ subsets: ["latin"] });

export default function Home() {
  return (
    <main>
      <ChakraProvider>
        <Container className="p-10">
          <div className="flex flex-row justify-center">
            <SiteTitle />
          </div>
          <div className="p-5 flex flex-col">
            <div className="p-2">
              <PasswordInput placeHolderText={"Secret"} />
            </div>
            <div className="p-2">
              <PasswordInput placeHolderText={"Secret Key"} />
            </div>
            <div className="p-2">
              <InputComponent inputText={"Bucket Name"} />
            </div>
            <div className="p-2">
              <InputFileComponent inputText={"Select File"} />
            </div>
            <div className="p-5">
              <ButtonComponent buttonText={"Upload"} />
            </div>
          </div>
        </Container>
      </ChakraProvider>
    </main>
  );
}
