import logo from "./logo.svg";
import "./App.css";
import { React } from "react";
import Button from "@material-ui/core/Button";
import FormControl from "@material-ui/core/FormControl";
import FormHelperText from "@material-ui/core/FormHelperText";
import InputLabel from "@material-ui/core/InputLabel";
import Input from "@material-ui/core/Input";
import Grid from "@material-ui/core/Grid";
import Typography from "@material-ui/core/Typography";

function App() {
  return (
    <>
      <link
        rel="stylesheet"
        href="https://fonts.googleapis.com/css?family=Roboto:300,400,500,700&display=swap"
      />
      <Grid container spacing={2}>
        <Grid item xs={6}>
          <Typography>Please enter a blog post URL</Typography>
        </Grid>
        <Grid item xs={3}>
          <form method="POST" action="/url">
            <FormControl>
              <InputLabel htmlFor="url">URL</InputLabel>
              <Input id="url" name="url" aria-describedby="url-description" />
              <FormHelperText id="url-description">
                Put your blog's URL here.
              </FormHelperText>
            </FormControl>
          </form>
        </Grid>
      </Grid>
    </>
  );
}

export default App;
