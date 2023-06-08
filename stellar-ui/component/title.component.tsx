import React from "react";
import { Text } from "@chakra-ui/react";
import { Sparkles } from "lucide-react";

export default function SiteTitle() {
  return (
    <div className="flex">
      <Sparkles />
      <Text fontSize="6xl">Stellar</Text>
      <Text fontSize="6xl">Upload</Text>
    </div>
  );
}
