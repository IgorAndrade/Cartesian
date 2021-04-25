package domain

import "testing"

func TestPoint_DistanceTo(t *testing.T) {
	type fields struct {
		X int
		Y int
	}
	type args struct {
		p2 Point
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		{
			name: "Distance 3 (positive values) ",
			fields: fields{
				X: 10,
				Y: 10,
			},
			args: args{
				p2: Point{
					X: 12,
					Y: 11,
				},
			},
			want: 3,
		},
		{
			name: "Distance 3 (negative values) ",
			fields: fields{
				X: 1,
				Y: 1,
			},
			args: args{
				p2: Point{
					X: 2,
					Y: -1,
				},
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Point{
				X: tt.fields.X,
				Y: tt.fields.Y,
			}
			if got := p.DistanceTo(tt.args.p2); got != tt.want {
				t.Errorf("Point.DistanceTo() = %v, want %v", got, tt.want)
			}
		})
	}
}
