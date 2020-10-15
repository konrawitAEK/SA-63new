import React, { useState, useEffect } from 'react';
import { makeStyles } from '@material-ui/core/styles';
import Table from '@material-ui/core/Table';
import TableBody from '@material-ui/core/TableBody';
import TableCell from '@material-ui/core/TableCell';
import TableContainer from '@material-ui/core/TableContainer';
import TableHead from '@material-ui/core/TableHead';
import TableRow from '@material-ui/core/TableRow';
import Paper from '@material-ui/core/Paper';
import Button from '@material-ui/core/Button';
import { DefaultApi } from '../../api/apis';
import { EntPositionassingment } from '../../api/models/EntPositionassingment';

const useStyles = makeStyles({
 table: {
   minWidth: 650,
 },
});
 
export default function ComponentsTable() {
 const classes = useStyles();
 const api = new DefaultApi();
 const [Positionassingments, setPositionassingments] = useState<EntPositionassingment[]>();
 const [loading, setLoading] = useState(true);
 
 useEffect(() => {
   const getPositionassingments = async () => {
     const res = await api.listPositionassingment({ limit: 10, offset: 0 });
     setLoading(false);
     setPositionassingments(res);
   };
   getPositionassingments();
 }, [loading]);
 
 const deletePositionassingments = async (id: number) => {
      const res = await api.deletePositionassingment({ id: id });
      setLoading(true);
 };
 
 return (
   <TableContainer component={Paper}>
     <Table className={classes.table} aria-label="simple table">
       <TableHead>
         <TableRow>
           <TableCell align="center">No.</TableCell>
           <TableCell align="center">Physician</TableCell>
           <TableCell align="center">Position</TableCell>
           <TableCell align="center">Department</TableCell>
           <TableCell align="center">DayStart</TableCell>
           <TableCell align="center">Manage</TableCell>
         </TableRow>
       </TableHead>
       <TableBody>
         {Positionassingments === undefined
         ? null
         : Positionassingments.map(item => (
           <TableRow key={item.id}>
             <TableCell align="center">{item.id}</TableCell>
             <TableCell align="center">{item.dayStart}</TableCell>
             <TableCell align="center">
               <Button
                 onClick={() => {
                   deletePositionassingments(item.id);
                 }}
                 style={{ marginLeft: 10 }}
                 variant="contained"
                 color="secondary"
               >
                 Delete
               </Button>
             </TableCell>
           </TableRow>
         ))}
       </TableBody>
     </Table>
   </TableContainer>
 );
}
