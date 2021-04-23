import { ApiProperty } from "@nestjs/swagger";

export class Types {
  @ApiProperty()
  types: Array<string>;
}

export class TypeCompatibility {
  @ApiProperty()
  type: string;

  @ApiProperty()
  compatibility: Array<number>;

  @ApiProperty()
  typeOrder: Types;
}

export class TypeCompatibilityTable {
  @ApiProperty({ type: [TypeCompatibility] })
  compatibilityTable: Array<TypeCompatibility>;

  @ApiProperty()
  typeOrder: Types;
}
