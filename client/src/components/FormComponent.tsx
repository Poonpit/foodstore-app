import React, { ChangeEvent, FormEvent, useState } from "react";
import {
  Box,
  Button,
  Checkbox,
  FormControlLabel,
  FormGroup,
  TextField,
  Typography,
} from "@mui/material";

interface FormValues {
  redSet: number;
  greenSet: number;
  blueSet: number;
  yellowSet: number;
  pinkSet: number;
  purpleSet: number;
  orangeSet: number;
  hasMemberCard: boolean;
}

interface FormComponentProps {
  onSubmit: (
    formValues: Omit<
      FormValues,
      | "redSet"
      | "greenSet"
      | "blueSet"
      | "yellowSet"
      | "pinkSet"
      | "purpleSet"
      | "orangeSet"
    > & { items: { [key: string]: number } }
  ) => void;
}

const FormComponent: React.FC<FormComponentProps> = ({ onSubmit }) => {
  const [formValues, setFormValues] = useState<FormValues>({
    redSet: 0,
    greenSet: 0,
    blueSet: 0,
    yellowSet: 0,
    pinkSet: 0,
    purpleSet: 0,
    orangeSet: 0,
    hasMemberCard: false,
  });

  const handleChange = (event: ChangeEvent<HTMLInputElement>) => {
    const { name, value, type, checked } = event.target;
    const parsedValue =
      type === "checkbox" ? checked : Math.max(0, Number(value));
    setFormValues({
      ...formValues,
      [name]: parsedValue,
    });
  };

  const formatKey = (key: string) => {
    // Capitalize the first letter of the key and keep "set" in lowercase
    return key.charAt(0).toUpperCase() + key.slice(1).replace("Set", " set");
  };

  const handleSubmit = (event: FormEvent<HTMLFormElement>) => {
    event.preventDefault();

    // Create items object with non-zero values
    const items = Object.fromEntries(
      Object.entries(formValues)
        .filter(([key, value]) => key !== "hasMemberCard" && value > 0)
        .map(([key, value]) => [formatKey(key), value])
    );

    // Create request data
    const requestData = {
      items,
      hasMemberCard: formValues.hasMemberCard,
    };

    onSubmit(requestData);
  };

  // Define an array of set names
  const setNames = [
    "red",
    "green",
    "blue",
    "yellow",
    "pink",
    "purple",
    "orange",
  ];

  return (
    <Box
      component="form"
      onSubmit={handleSubmit}
      sx={{
        display: "flex",
        flexDirection: "column",
        alignItems: "center",
        justifyContent: "center",
        maxWidth: 400,
        mx: "auto",
        mt: 4,
        p: 2,
        border: "1px solid #ccc",
        borderRadius: "8px",
      }}
    >
      <Typography variant="h5" gutterBottom>
        Food Store Calculator
      </Typography>
      <FormGroup>
        {setNames.map((setName) => (
          <TextField
            key={setName}
            label={`${setName.charAt(0).toUpperCase() + setName.slice(1)} set`}
            type="number"
            name={`${setName}Set`}
            value={formValues[`${setName}Set` as keyof FormValues]}
            onChange={handleChange}
            margin="normal"
            InputProps={{ inputProps: { min: 0 } }}
            fullWidth
          />
        ))}
        <FormControlLabel
          control={
            <Checkbox
              name="hasMemberCard"
              checked={formValues.hasMemberCard}
              onChange={handleChange}
            />
          }
          label="Do you have a member card?"
        />
        <Button
          variant="contained"
          color="primary"
          type="submit"
          sx={{ mt: 2 }}
        >
          Calculate Total
        </Button>
      </FormGroup>
    </Box>
  );
};

export default FormComponent;
