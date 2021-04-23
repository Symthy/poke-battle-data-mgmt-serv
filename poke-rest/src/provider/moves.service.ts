import { Move, Moves } from '@/models/moves';
import { Injectable } from '@nestjs/common';

@Injectable()
export class MovesService {
  getMove(id: string): Move {
    throw new Error('Method not implemented.');
  }

  getMoves(): Moves {
    return new Moves();
  }
}
