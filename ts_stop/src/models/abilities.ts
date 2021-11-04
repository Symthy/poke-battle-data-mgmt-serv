import { ApiProperty } from "@nestjs/swagger";

export class Ability {

  @ApiProperty()
  id: string;

  @ApiProperty()
  name: string;

  @ApiProperty()
  description: string;

  @ApiProperty()
  physicalMovePowerCorrectionValue: number;

  @ApiProperty()
  specialMovePowerCorrectionValue: number;

  @ApiProperty()
  attackPowerCorrectionValue: number;

  @ApiProperty()
  specialAttackPowerCorrectionValue: number;

  @ApiProperty()
  attackCorrection: number;

  @ApiProperty()
  specialAttackCorrection: number;

  @ApiProperty()
  damageCorrectionValue: number;

  @ApiProperty()
  weightCorrectionValue: number;
}

export class Abilities {
  @ApiProperty()
  abilities: Array<Ability>;

  @ApiProperty()
  count: number;
}

export class PokemonAbility {
  @ApiProperty()
  pokedexId: string;

  @ApiProperty()
  moves: Array<Ability>;
}
