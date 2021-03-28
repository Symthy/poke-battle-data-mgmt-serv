import { Injectable } from '@nestjs/common';
import { Move, Moves } from 'src/app/models/moves';

@Injectable()
export class MovesService {
  getMove(id: string): Move {
    throw new Error('Method not implemented.');
  }

  getMoves(): Moves {
    return new Moves();
  }
}
