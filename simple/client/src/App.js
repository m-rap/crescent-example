// import logo from './logo.svg';
import './App.css';
import { useEffect } from 'react';
import CrescentClient from 'crescent-client';

function App() {
  let symKey = null;
  let crescentClient = null;
  
  useEffect(() => {
    crescentClient = new CrescentClient();
  });

  let content;

  content = (
    <div>
      <p>Hello, Rian</p>
      <button onClick={() => {
        crescentClient = new CrescentClient();
      }}>reinstantiate crescentClient</button>
      <br />
      <button onClick={() => {
        crescentClient.post("http://localhost:8080/api/func1", { data: "data rahasia lho" });
      }}>call api</button>
    </div>
  );

  return (
    <div>
      <p>hello world from app</p>
      {content}
    </div>
  );
}

// function App() {
//   return (
//     <div className="App">
//       <header className="App-header">
//         <img src={logo} className="App-logo" alt="logo" />
//         <p>
//           Edit <code>src/App.js</code> and save to reload.
//         </p>
//         <a
//           className="App-link"
//           href="https://reactjs.org"
//           target="_blank"
//           rel="noopener noreferrer"
//         >
//           Learn React
//         </a>
//       </header>
//     </div>
//   );
// }

export default App;
