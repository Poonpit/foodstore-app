import React from "react";
import FormComponent from "./components/FormComponent";
import { Box, Container } from "@mui/material";

interface RequestData {
  items: { [key: string]: number };
  hasMemberCard: boolean;
}

const App: React.FC = () => {
  const handleSubmit = async (formValues: RequestData) => {
    try {
      console.log(formValues);
      const response = await fetch("http://localhost:3000/calculate", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify(formValues),
      });

      const responseText = await response.text(); // พead the response as text

      if (response.ok) {
        const result = JSON.parse(responseText); // Attempt to parse the response as JSON
        const totalFormatted = parseFloat(result.total).toFixed(2); // Format total to 2 decimal places
        alert(`Total: ${totalFormatted} THB`);
      } else {
        console.error("Failed to calculate total:", responseText);
      }
    } catch (error) {
      console.error("Error:", error);
    }
  };

  return (
    <Container
      maxWidth="sm"
      sx={{
        display: "flex",
        alignItems: "center",
        justifyContent: "center",
        minHeight: "100vh",
        bgcolor: "background.default",
      }}
    >
      <Box
        sx={{
          width: "100%",
          padding: 2,
          boxShadow: 3,
          borderRadius: 2,
          bgcolor: "background.paper",
        }}
      >
        <FormComponent onSubmit={handleSubmit} />
      </Box>
    </Container>
  );
};

export default App;
