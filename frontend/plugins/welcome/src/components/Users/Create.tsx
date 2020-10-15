import React, { useState , useEffect } from 'react';
import { Link as RouterLink } from 'react-router-dom';
import {
  Content,
  Header,
  Page,
  pageTheme,
  ContentHeader,
} from '@backstage/core';
import { makeStyles, Theme, createStyles } from '@material-ui/core/styles';
import TextField from '@material-ui/core/TextField';
import Button from '@material-ui/core/Button';
import FormControl from '@material-ui/core/FormControl';
import { Alert } from '@material-ui/lab';
import { DefaultApi } from '../../api/apis';

import SaveIcon from '@material-ui/icons/Save';
import { InputLabel, MenuItem, Select } from '@material-ui/core';
import CssBaseline from '@material-ui/core/CssBaseline';
import Typography from '@material-ui/core/Typography';
import Container from '@material-ui/core/Container';
import TableCell from '@material-ui/core/TableCell';
import Avatar from '@material-ui/core/Avatar';
import Box from '@material-ui/core/Box';
import Link from '@material-ui/core/Link';

import { EntPhysician } from '../../api/models/EntPhysician';
import { EntDepartment } from '../../api/models/EntDepartment';
import { EntPosition } from '../../api/models/EntPosition';

//import Swal from 'sweetalert2';

const useStyles = makeStyles((theme: Theme) =>
  createStyles({
    root: {
      display: 'flex',
      flexWrap: 'wrap',
      justifyContent: 'center',
    },
    margin: {
      margin: theme.spacing(2),
    },
    withoutLabel: {
      marginTop: theme.spacing(2),
    },
    textField: {
      width: '25ch',
    },
  }),
);

const initialUserState = {
  formdata: Number,
  formdepartment: Number,
  formposition: Number,
  DayStart: Date
};

export default function Create() {
  const classes = useStyles();
  const profile = { givenName: 'to PositionAssingment' };
  const api = new DefaultApi();

  const [Positionassingment, setPositionassingment] = React.useState(initialUserState);
  const [status, setStatus] = useState(false);
  const [alert, setAlert] = useState(true);

  const [physicians, setPhysicians] = React.useState<EntPhysician[]>([]);
  const [departments, setDepartments] = React.useState<EntDepartment[]>([]);
  const [positions, setPositions] = React.useState<EntPosition[]>([]);

  //const Toast = Swal.mixin({
  //  toast: true,
  //  position: 'top-end',
  //  showConfirmButton: false,
  //  timer: 3000,
  //  timerProgressBar: true,
  //  didOpen: toast => {
  //    toast.addEventListener('mouseenter', Swal.stopTimer);
  //    toast.addEventListener('mouseleave', Swal.resumeTimer);
  //  },
  //});

  const getPhysician = async () => {
    const res = await api.listPhysician({ limit: 10, offset: 0 });
    setPhysicians(res);
  };

  const getDepartment = async () => {
    const res = await api.listDepartment({ limit: 10, offset: 0 });
    setDepartments(res);
  };

  const getPosition = async () => {
    const res = await api.listPosition({ limit: 10, offset: 0 });
    setPositions(res);
  };

  useEffect(() => {
    getPhysician();
    getDepartment();
    getPosition();
  }, []);

  // set data to object Positionassingment
  const handleChange = (
    event: React.ChangeEvent<{ name?: string; value: unknown }>,
  ) => {
    const { value } = event.target;
    setPositionassingment({ ...Positionassingment, [name]: value });
    console.log(Positionassingment);
  };
  function clear() {
    setPositionassingment({});
  }
  // function save data
  function save() {
    const apiUrl = 'http://localhost:8080/api/v1/Positionassingment';
    const requestOptions = {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(Positionassingment),
    };

    console.log(Positionassingment); // log ดูข้อมูล สามารถ Inspect ดูข้อมูลได้ F12 เลือก Tab Console
    fetch(apiUrl, requestOptions)
      .then(response => response.json())
      .then(data => {
        console.log(data);
        if (data.status === true) {
          
          Toast.fire({
            icon: 'success',
            title: 'บันทึกข้อมูลสำเร็จ',
          });
        } else {
          Toast.fire({
           icon: 'error',
            title: 'บันทึกข้อมูลไม่สำเร็จ',
          });
        }
      });
  }

  return (
    <Page theme={pageTheme.home}>
      <Header
        title={`Welcome ${profile.givenName || 'to PositionAssingment'}`}
        subtitle="Select the Position you want to be in."
      >

        <Avatar>D</Avatar>
        <Typography component="div" variant="body1">
          <Box color="Dang@gmail.com">Dang@gmail.com</Box>
          <Box color="secondary.main"></Box>
          <Link
            component="button"
            variant="body2"
            onClick={() => {
              console.info("I'm a button.");
            }}
          >
            Logout
    </Link>
        </Typography>

      </Header>
      <Content>

        <div className={classes.root}>
          <form noValidate autoComplete="off">

            <TableCell align="left">
              <React.Fragment>
                <CssBaseline />
                <Container maxWidth="sm">
                  <Typography component="div" style={{ backgroundColor: '#cfe8fc', height: '50vh', width: '100vh' }} />

                </Container>
              </React.Fragment>
            </TableCell>

            <TableCell align="left">

              <FormControl
                fullWidth
                className={classes.margin}
                variant="outlined"
              >
                <TextField
                  id="datetime-local"
                  label="DayStart"
                  type="datetime-local"
                  defaultValue="2020-09-18T10:30"
                  className={classes.textField}
                  InputLabelProps={{
                    shrink: true,
                  }}
                  onChange={handleChange}
                />
              </FormControl>


              <FormControl
                fullWidth
                className={classes.margin}
                variant="outlined"
              >
                <InputLabel>เลือกผู้ใช้ระบบ</InputLabel>
                <Select
                  name="Physician"
                  value={Positionassingment.formdata || ''}
                  style={{ width: 300 }}
                  onChange={handleChange}
                >
                  {physicians.map(item => {
                    return (
                      <MenuItem key={item.id} value={item.id}>
                        {item.eMAIL}
                      </MenuItem>
                    );
                  })}
                </Select>
              </FormControl>

              <FormControl
                fullWidth
                className={classes.margin}
                variant="outlined"
              >
                <InputLabel>เลือกแผนก</InputLabel>
                <Select
                  name="Department"
                  value={Positionassingment.formdepartment || ''}
                  style={{ width: 300 }}
                  onChange={handleChange}
                >
                  {departments.map(item => {
                    return (
                      <MenuItem key={item.id} value={item.id}>
                        {item.departmentname}
                      </MenuItem>
                    );
                  })}
                </Select>
              </FormControl>

              <FormControl
                fullWidth
                className={classes.margin}
                variant="outlined"
              >
                <InputLabel>เลือกตำแหน่ง</InputLabel>
                <Select
                  name="Position"
                  value={Positionassingment.formposition || ''}
                  onChange={handleChange}
                  style={{ width: 300 }}
                >
                  {positions.map(item => {
                    return (
                      <MenuItem key={item.id} value={item.id}>
                        {item.nameposition}
                      </MenuItem>
                    );
                  })}
                </Select>
              </FormControl>
            </TableCell>
            <div className={classes.margin}>
              <TableCell align="right">
                <Button
                  variant="contained"
                  color="primary"
                  size="large"
                  startIcon={<SaveIcon />}
                  style={{ marginLeft: 632, width: 200 }}
                  onClick={save}
                >
                  SAVE DATA
              </Button>
              </TableCell>

              <TableCell align="right">
                <Button
                  style={{ marginLeft: 1 }}
                  component={RouterLink}
                  to="/"
                  variant="contained"
                >
                  Back
             </Button>
              </TableCell>

            </div>
          </form>
        </div>
      </Content>

    </Page>
  );
}
