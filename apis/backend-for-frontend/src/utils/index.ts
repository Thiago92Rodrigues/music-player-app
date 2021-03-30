import fs from 'fs';
import { promisify } from 'util';

export const getFileStatus = promisify(fs.stat);

export const delay = (seconds: number) => {
  return new Promise(resolve => setTimeout(resolve, seconds * 1000));
};

export const convertMonthToString = (month: number): string => {
  // prettier-ignore
  switch (month) {
    case 0: return 'Jan';
    case 1: return 'Feb';
    case 2: return 'Mar';
    case 3: return 'Apr';
    case 4: return 'May';
    case 5: return 'Jun';
    case 6: return 'Jul';
    case 7: return 'Aug';
    case 8: return 'Sep';
    case 9: return 'Oct';
    case 10: return 'Nov';
    case 11: return 'Dec';
    default: return '';
  }
};
