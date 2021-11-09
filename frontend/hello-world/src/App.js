import { useEffect, useState } from 'react';
import './App.css';
import { Box } from '@mui/system';
import { Typography } from '@mui/material';

function App() {
  const url = 'https://hw-be.herokuapp.com/';
  const [data, setData] = useState(null);

  const fetchData = async () => {
    try {
      const response = await fetch(url);
      return await response.json();
    } catch (error) {
      console.error(error);
    }
  }

  useEffect(() => {
    fetchData().then((res) => {
      setData(res)
    });
  }, [])

  return (
    <Box style={{ backgroundColor: '#173A5E', height: '100vh' }}>
      <Box style={{ textAlign: "center", display: "flex", flexDirection: "column" }}>
        <Box>
          <Typography color="white" variant="h4">Homework 1: Display Hello World</Typography>
        </Box>
        <Box>
          <Typography variant="h4" color="white">
            Hello World
          </Typography>
        </Box>
      </Box>
      <Box>
        <Typography color="white" variant="h4" style={{ textAlign: "center" }}>Homework 2: Make an API Call to render random string</Typography>
      </Box>
      <Box
        style={{
          border: "1px solid white",
          borderRadius: "10px",
          margin: "2rem",
          paddingTop: "2rem", 
          paddingBottom: "2rem", 
        }}
      >
        <Typography color="white" variant="h4" style={{ textAlign: "center" }}>{data}</Typography>
      </Box>
    </Box>
  );
}

export default App;