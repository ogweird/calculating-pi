#include <iostream>
#include <math.h>

class c_pi {
public:
	double leibniz(double n) {
		if (n == 0) {
			return n;
		}

		double pi = 0;
		double sign = 1;

		for (double i = 1; i <= (n * 2); i += 2) {
			pi = pi + sign * (4 / i);
			sign = -sign;
		}

		return pi;
	}

	double wallis(double n) {
		if (n == 0) {
			return n;
		}
		
		double pi = 4.;

		for (double i = 3; i <= (n + 2); i += 2) {
			pi = pi * ((i - 1) / i) * ((i + 1) / i);
		}

		return pi;
	}

	double nilakantha(double n) {
		if (n == 0) {
			return n;
		}

		double pi = 3;
		double sign = 1;

		for (double i = 2; i <= (n * 2); i += 2) {
			pi = pi + sign * (4 / (i * (i + 1) * (i + 2)));
			sign = -sign;
		}

		return pi;
	}
};

inline auto pi = new c_pi;