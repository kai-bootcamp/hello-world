import { Fragment, useEffect, useState } from 'react';
import './App.css';
import Table from '@mui/material/Table';
import TableBody from '@mui/material/TableBody';
import TableCell from '@mui/material/TableCell';
import TableContainer from '@mui/material/TableContainer';
import TableHead from '@mui/material/TableHead';
import TableRow from '@mui/material/TableRow';
import Paper from '@mui/material/Paper';
import { Box } from '@mui/system';
import { Checkbox, Typography } from '@mui/material';

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
      <Box style={{ textAlign: "center" }}>
        <Typography variant="h4" color="white">
          Hello World
        </Typography>
      </Box>
      {data && data.length > 0 ?
        <Box style={{ padding: "2rem" }}>
          <TableContainer component={Paper}>
            <Table sx={{ minWidth: 650 }} aria-label="simple table">
              <TableHead>
                <TableRow>
                  <TableCell>
                    <strong>Name</strong>
                  </TableCell>
                  <TableCell align="right">
                    <strong>Age</strong>
                  </TableCell>
                  <TableCell align="right">
                    <Box style={{
                      display: "flex",
                      flexDirection: "column"
                    }}>
                      <Box>
                        <strong>Sex</strong>
                      </Box>
                      <Box>
                        <Typography variant="caption">Checked is male</Typography>
                      </Box>
                    </Box>
                  </TableCell>
                  <TableCell align="right">
                    <strong>Salary</strong>
                  </TableCell>
                </TableRow>
              </TableHead>
              <TableBody>
                {data.map((item) => (
                  <TableRow
                    key={item.id}
                    sx={{ '&:last-child td, &:last-child th': { border: 0 } }}
                  >
                    <TableCell component="th" scope="row">
                      {item.name}
                    </TableCell>
                    <TableCell align="right">{item.age}</TableCell>
                    <TableCell align="right">
                      {+item.sex === 0 ? 
                        <Fragment>
                          <Checkbox checked disabled/>
                        </Fragment>
                      : 
                        <Fragment>
                          <Checkbox disabled/>
                        </Fragment>}
                    </TableCell>
                    <TableCell align="right">{item.salary}</TableCell>
                  </TableRow>
                ))}
              </TableBody>
            </Table>
          </TableContainer>
        </Box>
        : 
        null
      }
    </Box>
  );
}

export default App;