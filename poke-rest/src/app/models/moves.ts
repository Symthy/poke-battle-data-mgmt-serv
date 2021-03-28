import { ApiProperty } from "@nestjs/swagger";

export class Move {

  @ApiProperty()
  id: string;

  @ApiProperty()
  name: string;

  @ApiProperty()
  type: string;

  @ApiProperty()
  moveSpecies: string;

  @ApiProperty()
  power: number;

  @ApiProperty()
  accuracy: number;

  @ApiProperty()
  pp: number;

  @ApiProperty()
  isContactMove: boolean;

  @ApiProperty()
  isCanGuard: boolean;

}

export class Moves {

  @ApiProperty()
  moves: Array<Move>;

  @ApiProperty()
  count: number;

}

export class PokemonMoves {
  @ApiProperty()
  pokedexId: string;

  @ApiProperty()
  moves: Array<Move>;
}
