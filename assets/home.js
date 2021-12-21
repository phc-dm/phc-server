
function pointInInterval(a, b, x) {
  return a <= x && x <= b;
}

function intervalIntersectsInterval(a, b, c, d) {
  if (c <= a) {
    return a <= d;
  }

  return pointInInterval(a, b, c);
}

function rectIntersectRect(x1, y1, x2, y2, x3, y3, x4, y4) {
  return intervalIntersectsInterval(x1, x2, x3, x4) && intervalIntersectsInterval(y1, y2, y3, y4);
}

function intersects(x1, y1, x2, y2, x3, y3, x4, y4) {
  const a_minx = Math.min(x1, x2);
  const a_maxx = Math.max(x1, x2);
  const a_miny = Math.min(y1, y2);
  const a_maxy = Math.max(y1, y2);

  const b_minx = Math.min(x3, x4);
  const b_maxx = Math.max(x3, x4);
  const b_miny = Math.min(y3, y4);
  const b_maxy = Math.max(y3, y4);

  if (!rectIntersectRect(a_minx, a_miny, a_maxx, a_maxy, b_minx, b_miny, b_maxx, b_maxy))
    return false;

  const a_dx = x2 - x1;
  const a_dy = y2 - y1;

  const b_dx = x4 - x3;
  const b_dy = y4 - y3;

  const det = -b_dx * a_dy + a_dx * b_dy;

  if (Math.abs(det) === 0 && -a_dy * (x3 - x1) + a_dx * (y3 - y1) === 0) {
    return y1 <= y3 <= y2 || y1 <= y4 <= y2;
  }

  const s = (-a_dy * (x1 - x3) + a_dx * (y1 - y3)) / det;
  const t = (+b_dx * (y1 - y3) - b_dy * (x1 - x3)) / det;

  return s >= 0 && s <= 1 && t >= 0 && t <= 1;
}

const $canvas = document.querySelector('#circuit-pattern');
let g = $canvas.getContext('2d');

let WIDTH = $canvas.offsetWidth;
let HEIGHT = $canvas.offsetHeight;

const TS = 20;

let ROWS = HEIGHT / TS;
let COLS = WIDTH / TS;

function updateCanvasDimensions() {
  g = $canvas.getContext('2d');

  WIDTH = $canvas.offsetWidth;
  HEIGHT = $canvas.offsetHeight;

  ROWS = HEIGHT / TS;
  COLS = WIDTH / TS;

  $canvas.width = WIDTH * devicePixelRatio;
  $canvas.height = HEIGHT * devicePixelRatio;
  g.scale(devicePixelRatio, devicePixelRatio);
}

window.addEventListener('resize', () => updateCanvasDimensions());
updateCanvasDimensions();

const state = {
  wires: [], // :: [(Int, Int)]
  nextWire: null,
  nextWireIndex: 0,
  nextSegmentLen: 0,
  t: 0,
};

function randomFloat(min, max) {
  return Math.random() * (max - min) + min;
}

function randomInt(min, max) {
  return Math.floor(Math.random() * (max - min + 1) + min);
}

function randomChoice(list) {
  return list[randomInt(0, list.length - 1)];
}

const DIRS = {
  left: [-1, 1],
  center: [0, 1],
  right: [+1, 1],
};

const NEXT_DIRS = {
  left: ['center'],
  center: ['left', 'right'],
  right: ['center'],
};

function canPlaceSegment([[p1x, p1y], [p2x, p2y]]) {
  return state.wires.every(wire => {
    return everyWireSegment(wire, ([[sx, sy], [lx, ly]]) => {
      return !intersects(sx, sy, lx, ly, p1x, p1y, p2x, p2y);
    });
  })
}

function everyWireSegment(wire, segmentFn) {
  let [start, ...rest] = wire;

  return rest.every((pt) => {
    const r = segmentFn([start, pt]);
    start = pt;
    return r;
  });
}

function generateWire() {
  const pieceCount = Math.pow(randomFloat(1, 2.5), 2) | 0;
  const startPoint = [
    randomInt(0, COLS),
    Math.pow(randomFloat(0, ROWS / 12), 2) | 0,
  ];
  const sizes = Array(pieceCount).fill(0).map(() => randomInt(2, 4) | 0);

  const dirs = [randomChoice(Object.keys(DIRS))];
  for (let i = 0; i < pieceCount - 1; i++) {
    const last = dirs[dirs.length - 1];
    dirs.push(randomChoice(NEXT_DIRS[last]));
  }

  const wire = [startPoint];
  for (let i = 0; i < pieceCount; i++) {
    const [lpx, lpy] = wire[wire.length - 1];
    const [dx, dy] = DIRS[dirs[i]];
    const s = sizes[i];

    const s_bonus = (dirs[i] === 'center' ? s * 2 : s) | 0;

    wire.push([lpx + dx * s_bonus, lpy + dy * s_bonus]);
  }

  const ok = everyWireSegment(wire, segment => {
    return canPlaceSegment(segment);
  });

  if (!ok)
    return null;

  return wire;
}

// // LinearintERPolation
function lerp(a, b, t) {
  return t * (b - a) + a;
}
function lerpPoint([x1, y1], [x2, y2], t) {
  return [lerp(x1, x2, t), lerp(y1, y2, t)];
}

const CELLS_PER_SECOND = 15;

let lastTime = new Date().getTime();

function update() {
  const now = new Date().getTime();
  const delta = (now - lastTime) / 1000;
  lastTime = now;

  if (state.nextWire) {
    state.t += CELLS_PER_SECOND * delta;

    if (state.t >= 1.0) {
      const [[x1, y1], [x2, y2]] = state.nextWire.slice(state.nextWireIndex - 2, state.nextWireIndex);
      const dx = Math.abs(x2 - x1);
      const dy = Math.abs(y2 - y1);
      state.t = 0;
      state.nextSegmentLen = Math.sqrt(dx * dx + dy * dy);

      state.nextWireIndex++;
    }

    if (state.nextWireIndex > state.nextWire.length) {
      state.wires.push(state.nextWire);
      state.nextWire = null;
    }
  } else {
    const wire = generateWire();
    if (wire) {
      state.nextWire = wire;
      state.nextWireIndex = 2;

      const [[x1, y1], [x2, y2]] = state.nextWire.slice(state.nextWireIndex - 2, state.nextWireIndex);
      const dx = Math.abs(x2 - x1);
      const dy = Math.abs(y2 - y1);
      state.t = 0;
      state.nextSegmentLen = Math.sqrt(dx * dx + dy * dy);
    }
  }
}

function getColors() {
  if (document.body.classList.contains('dark-mode')) {
    return {
      BACKGROUND: '#282828',
      // CIRCUIT: '#202020',
      CIRCUIT: '#38302e',
    }
  } else {
    return {
      BACKGROUND: '#eaeaea',
      CIRCUIT: '#d4d4d4',
    };
  }
}

function render() {
  const { BACKGROUND, CIRCUIT } = getColors();

  g.clearRect(0, 0, WIDTH, HEIGHT);

  g.strokeStyle = CIRCUIT;
  g.lineWidth = 3;

  state.wires.forEach(wire => {
    const [[sx, sy], ...rest] = wire;

    function renderDot(x, y) {
      if (y === 0) {
        return;
      }

      g.beginPath();
      g.ellipse(x * TS, y * TS, TS / 5, TS / 5, 0, 0, Math.PI * 2);
      g.stroke();
    }

    renderDot(sx, sy);

    g.beginPath();
    g.moveTo(sx * TS, sy * TS);
    rest.forEach(([x, y]) => {
      g.lineTo(x * TS, y * TS);
    });
    g.stroke();

    const [lx, ly] = wire[wire.length - 1];

    renderDot(lx, ly);

    [wire[0], wire[wire.length - 1]].forEach(([x, y]) => {
      if (y === 0) {
        return;
      }

      g.fillStyle = BACKGROUND;
      g.beginPath();
      g.ellipse(x * TS, y * TS, TS / 5 - 1, TS / 5 - 1, 0, 0, Math.PI * 2);
      g.fill();
    });
  });

  if (state.nextWire) {
    const wirePart = state.nextWire.slice(0, state.nextWireIndex);
    const wire = [
      ...wirePart.slice(0, wirePart.length - 1),
      lerpPoint(
        wirePart[wirePart.length - 2],
        wirePart[wirePart.length - 1],
        state.t
      )
    ];
    const [[sx, sy], ...rest] = wire;

    function renderDot(x, y) {
      if (y === 0) {
        return;
      }

      g.beginPath();
      g.ellipse(x * TS, y * TS, TS / 5, TS / 5, 0, 0, Math.PI * 2);
      g.stroke();
    }

    renderDot(sx, sy);

    g.beginPath();
    g.moveTo(sx * TS, sy * TS);
    rest.forEach(([x, y]) => {
      g.lineTo(x * TS, y * TS);
    });
    g.stroke();

    const [lx, ly] = wire[wire.length - 1];

    renderDot(lx, ly);

    [wire[0], wire[wire.length - 1]].forEach(([x, y]) => {
      if (y === 0) {
        return;
      }

      g.fillStyle = BACKGROUND;
      g.beginPath();
      g.ellipse(x * TS, y * TS, TS / 5 - 1, TS / 5 - 1, 0, 0, Math.PI * 2);
      g.fill();
    });
  }

  requestAnimationFrame(render);
}


// g.translate(0.5, 0.5);

window.addEventListener('DOMContentLoaded', () => {
  // $canvas.classList.remove('hide');

  render();

  let i = 20;
  while (i > 0) {
    const wire = generateWire();

    if (wire) {
      state.wires.push(wire);
      i--;
    }
  }

  setInterval(() => {
    update();
  }, 1000 / 30);
})