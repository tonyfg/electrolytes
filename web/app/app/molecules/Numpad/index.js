import style from './style.scss';

import { linkEvent } from 'inferno';

export default function Numpad({ onClick }) {
  const clickHandler = linkEvent(onClick, handleClick);
  return (
    <table className={style.numpad} onMouseDown={clickHandler}>
      <tr>{[1, 2, 3].map(Btn)}</tr>
      <tr>{[4, 5, 6].map(Btn)}</tr>
      <tr>{[7, 8, 9].map(Btn)}</tr>
    </table>
  );
}

function Btn(n) {
  return (
    <td>
      <button value={n} className="rounded touch btn">
        {n}
      </button>
    </td>
  );
}

function handleClick(onClick, event) {
  event.preventDefault();
  event.stopPropagation();

  if (event.target.value != null) {
    // Only the buttons have values so we ignore everything else
    onClick(event.target.value);
  }
}
